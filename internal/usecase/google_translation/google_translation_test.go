package google_translation

import (
	"context"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestTranslate(t *testing.T) {
	ctx := context.Background()
	translator := NewTranslator(client)

	_, err := translator.Translate(ctx, language.Und, language.Und, "")
	assert.NotNil(t, err)

	tr, err := translator.Translate(ctx, language.Und, language.English, "")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "",
		DetectedLanguage: language.Und,
	}, tr)
	assert.Nil(t, err)

	tr, err = translator.Translate(ctx, language.Und, language.English, "Привет")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "Hello",
		DetectedLanguage: language.Russian,
	}, tr)
	assert.Nil(t, err)

	tr, err = translator.Translate(ctx, language.English, language.English, "Привет")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "Привет",
		DetectedLanguage: language.English,
	}, tr)
	assert.Nil(t, err)

	tr, err = translator.Translate(ctx, language.Russian, language.English, "Привет")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "Hello",
		DetectedLanguage: language.Russian,
	}, tr)
	assert.Nil(t, err)

	tr, err = translator.Translate(ctx, language.Und, language.English, "Привет")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "Hello",
		DetectedLanguage: language.Russian,
	}, tr)
	assert.Nil(t, err)
}
