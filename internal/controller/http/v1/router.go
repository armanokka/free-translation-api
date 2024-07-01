package v1

import (
	"github.com/armanokka/transloapi/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	log             usecase.Logger
	translator      usecase.TranslationService
	translationRepo usecase.TranslationRepo
	tracer          trace.Tracer
}

func NewRouter(handler *gin.RouterGroup, log usecase.Logger, translator usecase.TranslationService, translationRepo usecase.TranslationRepo, tracer trace.Tracer) {
	app := &App{
		log:             log,
		translator:      translator,
		translationRepo: translationRepo,
		tracer:          tracer,
	}
	handler.POST("/translate", app.TranslateHandler)
	handler.POST("/detect", app.DetectHandler)
}
