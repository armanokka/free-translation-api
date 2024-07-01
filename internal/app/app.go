package app

import (
	"context"
	"errors"
	"github.com/armanokka/transloapi/config"
	_ "github.com/armanokka/transloapi/docs"
	v1 "github.com/armanokka/transloapi/internal/controller/http/v1"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/armanokka/transloapi/internal/usecase/google_translation"
	"github.com/armanokka/transloapi/internal/usecase/redis_repo"
	"github.com/armanokka/transloapi/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT)
	defer signal.Stop(c)
	go func() {
		<-c
		cancel()
	}()

	log := logger.New(cfg.LogLevel)
	if cfg.LogLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	// Creating tracer
	shutdownTracerProvider, err := initProvider(ctx, cfg.OTLPGRPCReceiverDSN)
	if err != nil {
		return err
	}
	defer func() {
		if err = shutdownTracerProvider(ctx); err != nil {
			log.Error("shutdown", zap.Error(err))
		}
	}()
	tracer := otel.Tracer("transloapi")
	log.Info("Connected to OTEL collector")

	translator := google_translation.NewTranslator(client)

	// Creating redis translation repo
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + strconv.Itoa(cfg.RedisPort),
		Password: cfg.RedisPassword, // no password set
		DB:       0,                 // use default DB
	})
	if err = rdb.Ping(ctx).Err(); err != nil {
		return err
	}
	log.Info("Connected to Redis service")
	translationRepo := redis_repo.NewTranslationRepo(rdb, tracer)

	// Creating gin instance
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/uptime", UptimeHandler())
	router.GET("/languages-en.json", LanguagesListHandler(entity.SupportedLanguageCodesEn))
	router.GET("/languages-ru.json", LanguagesListHandler(entity.SupportedLanguageCodesRu))
	router.GET("/languages-zh.json", LanguagesListHandler(entity.SupportedLanguageCodesZh))

	v1.NewRouter(router.Group("/api/v1"), log, translator, translationRepo, tracer)

	// Handling signals. There is no graceful shutdown
	server := &http.Server{
		Addr:              "0.0.0.0:" + strconv.Itoa(cfg.Port),
		ReadHeaderTimeout: 30 * time.Second,
		Handler:           router,
	}
	if err = http2.ConfigureServer(server, nil); err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		err = server.Shutdown(ctx)
		if err != nil {
			log.Error("shutdown", zap.Error(err))
		}
	}()

	log.Info("Server launched", zap.Int("port", cfg.Port))

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
