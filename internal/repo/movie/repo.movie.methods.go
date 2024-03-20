package account

import (
	"context"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (r *Repo) ListMovie(ctx context.Context, size int32, token string) ([]domain.Movie, string, error) {
	ctx, span := tracer.Start(ctx, "repo.ListMovie")
	defer span.End()

	return []domain.Movie{{ID: 1}}, "", nil
}
