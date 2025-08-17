package usecase

import (
	"backend/src/url/entity"
	"context"
)

type URLUsecase interface {
	CleanupURL(ctx context.Context, input entity.URL) (entity.URL, error)
}
