package v1

import (
	"fmt"
)

// Response is embedded struct in every API's answer
type Response struct {
	Ok  bool   `json:"ok"`
	Err string `json:"error,omitempty"`
} // @name Response

type ResponseTranslation struct {
	Response
	TextLang       string   `json:"text_lang"`
	TranslatedText []string `json:"translated_text,omitempty"` // string if only one text is translated, array if more
} // @name ResponseTranslation

type ResponseDetection struct {
	Response
	DetectedLanguages []string `json:"detected_languages"`
}

func ErrInvalid(fieldName string) Response {
	return Response{Err: fmt.Sprintf("%s is invalid", fieldName)}
}

var ErrInternal = Response{Err: "internal server error"}
