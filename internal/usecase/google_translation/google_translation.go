package google_translation

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/armanokka/transloapi/internal/entity"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Translator struct {
	client *http.Client
}

type GoogleTranslateAPIResponse struct {
	Sentences []Sentence `json:"sentences"`
	Src       string     `json:"src"`
}

type Sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}

func NewTranslator(client *http.Client) Translator {
	return Translator{client: client}
}

func (c Translator) Translate(ctx context.Context, from language.Tag, to language.Tag, text string) (entity.Translation, error) {
	if to == language.Und {
		return entity.Translation{}, fmt.Errorf("empty \"from\" language")
	}
	if text == "" {
		return entity.Translation{
			TranslatedText:   "",
			DetectedLanguage: language.Und,
		}, nil
	}
	if from == to {
		return entity.Translation{
			TranslatedText:   text,
			DetectedLanguage: from,
		}, nil
	}
	values := url.Values{}
	values.Set("client", "gtx")
	values.Set("dt", "t")
	if from == language.Und {
		values.Set("sl", "auto")
	} else {
		values.Set("sl", from.String())
	}
	values.Set("tl", to.String())
	values.Set("q", text)
	values.Set("dj", "1")

	req, err := http.NewRequestWithContext(ctx, "GET", "https://translate.googleapis.com/translate_a/single?"+values.Encode(), nil)
	if err != nil {
		return entity.Translation{}, err
	}
	req.Header["Accept"] = []string{"*/*"}
	req.Header["Origin"] = []string{"https://www.google.com"}
	req.Header["Referer"] = []string{"https://www.google.com/"}
	req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"}

	resp, err := c.client.Do(req)
	if err != nil {
		return entity.Translation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.Translation{}, err
	}

	if resp.StatusCode != 200 {
		return entity.Translation{}, fmt.Errorf("translate: not 200 code: %d %s", resp.StatusCode, string(body))
	}

	var out GoogleTranslateAPIResponse
	if err = json.Unmarshal(body, &out); err != nil {
		return entity.Translation{}, err
	}

	translatedText := new(strings.Builder)
	for _, v := range out.Sentences {
		translatedText.WriteString(v.Trans)
	}
	srcLang, err := language.Parse(out.Src)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("translate: invalid src lang: %s", srcLang)
	}

	return entity.Translation{
		TranslatedText:   translatedText.String(),
		DetectedLanguage: srcLang,
	}, nil
}
