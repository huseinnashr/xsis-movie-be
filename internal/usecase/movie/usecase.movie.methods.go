package account

import (
	"context"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (u *Usecase) ListMovie(ctx context.Context, pageSize int32, pageToken string) ([]domain.Movie, string, error) {
	ctx, span := tracer.Start(ctx, "usecase.ListMovie")
	defer span.End()

	return u.movieRepo.ListMovie(ctx, pageSize, pageToken)
}

func (u *Usecase) CreateMovie(ctx context.Context, param domain.CreateMovieParam) (int64, error) {
	ctx, span := tracer.Start(ctx, "usecase.CreateMovie")
	defer span.End()

	return u.movieRepo.CreateMovie(ctx, param)
}

func (u *Usecase) GetMovie(ctx context.Context, id int64) (domain.Movie, error) {
	ctx, span := tracer.Start(ctx, "usecase.GetMovie")
	defer span.End()

	return u.movieRepo.GetMovie(ctx, id)
}

func (u *Usecase) UpdateMovie(ctx context.Context, param domain.UpdateMovieParam) error {
	ctx, span := tracer.Start(ctx, "usecase.UpdateMovie")
	defer span.End()

	return u.movieRepo.UpdateMovie(ctx, param)
}

func (u *Usecase) DeleteMovie(ctx context.Context, id int64) error {
	ctx, span := tracer.Start(ctx, "usecase.DeleteMovie")
	defer span.End()

	return u.movieRepo.DeleteMovie(ctx, id)
}
