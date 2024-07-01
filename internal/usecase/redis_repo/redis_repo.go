package redis_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/armanokka/transloapi/internal/entity"
	"github.com/go-redis/redis/v9"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/language"
	"strings"
)

type TranslationRepo struct {
	client *redis.Client
	tracer trace.Tracer
}

func NewTranslationRepo(client *redis.Client, tracer trace.Tracer) TranslationRepo {
	return TranslationRepo{client: client, tracer: tracer}
}

func (r TranslationRepo) SaveTranslation(ctx context.Context, from, to language.Tag, original, translation string) error {
	if from == language.Und || to == language.Und {
		return fmt.Errorf("SaveTranslation: unexpected language.Und")
	}

	ctx, span := r.tracer.Start(ctx, "SaveTranslation: redis HSET")
	defer span.End()

	return r.client.HSet(ctx, to.String(), original, from.String()+","+translation).Err()
}

func (r TranslationRepo) GetTranslation(ctx context.Context, to language.Tag, originalText string) (translation entity.Translation, err error) {
	if to == language.Und {
		return entity.Translation{}, fmt.Errorf("GetTranslation: passed language.Und")
	}

	ctx, span := r.tracer.Start(ctx, "SaveTranslation: redis HGET")
	defer span.End()

	cmd := r.client.HGet(ctx, to.String(), originalText)
	if cmd == nil {
		return entity.Translation{}, entity.ErrNotFound
	}
	if cmd.Err() != nil {
		if errors.Is(cmd.Err(), redis.Nil) {
			return entity.Translation{}, entity.ErrNotFound
		}
		return entity.Translation{}, cmd.Err()
	}
	value := cmd.Val()

	i := strings.Index(value, ",")
	if i == -1 {
		return entity.Translation{}, fmt.Errorf("index: ',' not found in cache")
	}
	from, err := language.Parse(value[:i])
	if err != nil {
		return entity.Translation{}, fmt.Errorf("GetTranslation: invalid from in cache: %v (toLang: %s, originalText: %s)", from, to, originalText)
	}
	return entity.Translation{
		TranslatedText:   value[i+1:],
		DetectedLanguage: from,
	}, nil
}

func (r TranslationRepo) DeleteTranslation(ctx context.Context, toLang language.Tag, originalText string) error {
	if toLang == language.Und {
		return fmt.Errorf("DeleteTranslation: passed language.Und")
	}

	ctx, span := r.tracer.Start(ctx, "SaveTranslation: redis HDEL")
	defer span.End()

	cmd := r.client.HDel(ctx, toLang.String(), originalText)
	if cmd != nil {
		return cmd.Err()
	}
	return fmt.Errorf("nil")
}
