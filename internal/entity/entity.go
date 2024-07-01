package entity

import (
	"fmt"
	"golang.org/x/text/language"
)

type Translation struct {
	TranslatedText   string
	DetectedLanguage language.Tag
}

var ErrNotFound = fmt.Errorf("not found")
