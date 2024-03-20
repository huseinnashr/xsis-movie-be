package account

import (
	"context"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (u *Usecase) ListMovie(ctx context.Context, size int32, token string) ([]domain.Movie, string, error) {
	ctx, span := tracer.Start(ctx, "usecase.Signup")
	defer span.End()

	return u.movieRepo.ListMovie(ctx, size, token)
}
