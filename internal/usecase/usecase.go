package usecase

import (
	"context"
	"github.com/armanokka/transloapi/internal/entity"
	"golang.org/x/text/language"
)

type TranslationService interface {
	Translate(ctx context.Context, from language.Tag, to language.Tag, text string) (entity.Translation, error)
}

type TranslationRepo interface {
	GetTranslation(ctx context.Context, to language.Tag, originalText string) (translation entity.Translation, err error)
	SaveTranslation(ctx context.Context, from, to language.Tag, original, translation string) error
	DeleteTranslation(ctx context.Context, toLang language.Tag, originalText string) error
}

type Logger interface {
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
	With(args ...interface{}) Logger
}
