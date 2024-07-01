package v1

import (
	"errors"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
	"strings"
	"sync"
)

// TranslateHandler godoc
// @Summary      Translate text
// @Description  Translate text and detect its language
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        text formData string true "Text to translate"
// @Param        from formData string false "Source language"
// @Param        to   formData string true "Target language"
// @Success      200  {object}  ResponseTranslation
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /translate [post]
func (app App) TranslateHandler(c *gin.Context) {
	log := app.log.With(zap.Strings("text", c.PostFormArray("text")),
		zap.String("to", c.PostForm("to")), zap.String("from", c.PostForm("from")))
	log.Debug("new request POST /v1/translate")

	ctx, span := app.tracer.Start(c, "POST /v1/translate")
	defer span.End()

	// Parsing texts
	texts := c.PostFormArray("text")
	if len(texts) == 0 {
		c.IndentedJSON(400, ErrInvalid("text"))
		return
	}
	log = log.With(zap.Strings("text", texts))

	// Parsing target language
	to, err := language.Parse(strings.TrimSpace(c.PostForm("to")))
	if err != nil {
		c.IndentedJSON(400, ErrInvalid("to"))
		return
	}
	log = log.With(zap.String("to", to.String()))

	// Parsing source language
	from, err := language.Parse(c.PostForm("from"))
	if err != nil {
		from = language.Und
	}
	log = log.With(zap.String("from", from.String()))

	// Translating
	var g errgroup.Group
	var mu sync.Mutex
	for i, original := range texts {
		i, original := i, original
		g.Go(func() error {
			// Searching for translation in cache
			translation, err := app.translationRepo.GetTranslation(ctx, to, original)
			if err != nil {
				if !errors.Is(err, entity.ErrNotFound) {
					log.Error("translationRepo.GetTranslation error", zap.String("original", original), zap.Error(err))
				}

				// If translation is not present in cache, translating manually
				translation, err = app.translator.Translate(ctx, from, to, original)
				if err != nil {
					log.Error("translate error", zap.Error(err))
					return err
				}

				if err = app.translationRepo.SaveTranslation(ctx, translation.DetectedLanguage, to, original, translation.TranslatedText); err != nil {
					log.Error("translationRepo.SaveTranslation", zap.Error(err))
				}

			}

			mu.Lock()
			defer mu.Unlock()
			texts[i] = translation.TranslatedText
			if from == language.Und {
				from = translation.DetectedLanguage
			}
			return nil
		})
	}
	if err = g.Wait(); err != nil {
		c.PureJSON(400, ErrInternal)
		log.Error("wait", zap.Error(err))
		return
	}

	// Responding
	c.PureJSON(200, ResponseTranslation{
		Response: Response{
			Ok: true,
		},
		TextLang:       from.String(),
		TranslatedText: texts,
	})
}
