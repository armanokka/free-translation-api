package v1

import (
	"errors"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
	"sync"
)

// DetectHandler godoc
// @Summary      Detect language the text is written in
// @Description  Detect language the text is written in
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        text formData string true "Text to translate"
// @Success      200  {object}  ResponseDetection
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /detect [post]
func (app App) DetectHandler(c *gin.Context) {
	log := app.log.With(zap.Strings("text", c.PostFormArray("text")))
	log.Debug("new request GET /v1/detect")

	ctx, span := app.tracer.Start(c, "POST /v1/detect")
	defer span.End()

	// Parsing text
	texts := c.PostFormArray("text")
	if len(texts) == 0 {
		c.AbortWithStatusJSON(400, ErrInvalid("text"))
		return
	}
	log = log.With(zap.Strings("texts", texts))
	span.SetAttributes(attribute.StringSlice("texts", texts))

	// Detecting language
	var g errgroup.Group
	var mu sync.Mutex
	for i, text := range texts {
		text, i := text, i
		g.Go(func() error {
			// Searching for translation in cache
			translation, err := app.translationRepo.GetTranslation(ctx, language.English, text)
			if err != nil {
				if !errors.Is(err, entity.ErrNotFound) {
					log.Error("translationRepo.GetTranslation error", zap.String("text", text))
				}

				// If translation is not present in cache, translating manually
				translation, err = app.translator.Translate(ctx, language.Und, language.English, text)
				if err != nil {
					log.Error("translate error", zap.Error(err))
					return err
				}

				// Caching results
				if err = app.translationRepo.SaveTranslation(ctx, translation.DetectedLanguage, language.English, text, translation.TranslatedText); err != nil {
					log.Error("translationRepo.SaveTranslation", zap.Error(err))
				}
			}

			mu.Lock()
			defer mu.Unlock()
			texts[i] = translation.DetectedLanguage.String()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Error("can't translate text", zap.Error(err))
		c.AbortWithStatusJSON(500, ErrInternal)
		return
	}

	log = log.With(zap.Strings("detected_languages", texts))
	span.SetStatus(codes.Ok, "responded")

	c.PureJSON(200, ResponseDetection{
		Response: Response{
			Ok: true,
		},
		DetectedLanguages: texts,
	})
}
