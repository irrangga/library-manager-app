package usecase

import (
	"backend/internal/url/entity"
	"context"
)

type URLUsecase interface {
	CleanupURL(ctx context.Context, input entity.URL) (entity.URL, error)
}
