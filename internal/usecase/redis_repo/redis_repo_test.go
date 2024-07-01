package redis_repo

import (
	"context"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"golang.org/x/text/language"
	"log"
	"testing"
)

func SetupNoopTracer() trace.Tracer {
	return noop.NewTracerProvider().Tracer("transloapi")
}

// SetupRedis launches local Redis instance via testcontainers. Returned testcontainers.Container MUST be terminated
func SetupRedis(ctx context.Context) (testcontainers.Container, *redis.Client) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start redis: %s", err)
	}
	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	return redisC, redis.NewClient(&redis.Options{
		Addr: endpoint,
	})
}

func TestTranslationRepo_GetTranslation(t *testing.T) {
	ctx := context.Background()

	redisC, rdb := SetupRedis(ctx)
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	repo := NewTranslationRepo(rdb, SetupNoopTracer())

	_, err := repo.GetTranslation(ctx, language.Und, "hello")
	assert.NotNil(t, err)

	_, err = repo.GetTranslation(ctx, language.English, "")
	assert.Error(t, err, entity.ErrNotFound)

	_, err = repo.GetTranslation(ctx, language.English, "hello")
	assert.Error(t, err, entity.ErrNotFound)

	assert.Nil(t, repo.SaveTranslation(ctx, language.Russian, language.English, "привет", "hello"))

	tr, err := repo.GetTranslation(ctx, language.English, "привет")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "hello",
		DetectedLanguage: language.Russian,
	}, tr)
	assert.Nil(t, err)
}

func TestTranslationRepo_SaveTranslation(t *testing.T) {
	ctx := context.Background()

	redisC, rdb := SetupRedis(ctx)
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	repo := NewTranslationRepo(rdb, SetupNoopTracer())

	assert.NotNil(t, repo.SaveTranslation(ctx, language.Und, language.English, "привет", "hello"))
	assert.NotNil(t, repo.SaveTranslation(ctx, language.English, language.Und, "привет", "hello"))
	assert.Nil(t, repo.SaveTranslation(ctx, language.Russian, language.English, "человек", "human"))
	tr, err := repo.GetTranslation(ctx, language.English, "человек")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "human",
		DetectedLanguage: language.Russian,
	}, tr)
	assert.Nil(t, err)
}

func TestTranslationRepo_DeleteTranslation(t *testing.T) {
	ctx := context.Background()

	redisC, rdb := SetupRedis(ctx)
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	repo := NewTranslationRepo(rdb, SetupNoopTracer())

	assert.Nil(t, repo.SaveTranslation(ctx, language.English, language.Russian, "programming", "программирование"))

	tr, err := repo.GetTranslation(ctx, language.Russian, "programming")
	assert.Equal(t, entity.Translation{
		TranslatedText:   "программирование",
		DetectedLanguage: language.English,
	}, tr)
	assert.Nil(t, err)

	assert.Nil(t, repo.DeleteTranslation(ctx, language.Russian, "programming"))

	_, err = repo.GetTranslation(ctx, language.Russian, "programming")
	assert.Equal(t, err, entity.ErrNotFound)
}
