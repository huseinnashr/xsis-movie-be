package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/huseinnashr/xsis-movie-be/api/v1"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (h *Handler) ListMovies(ctx context.Context, req *v1.ListMovieRequest) (*v1.ListMovieResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.ListMovie")
	defer span.End()

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(http.StatusText(http.StatusBadRequest), err.Error())
	}

	movies, nextPageToken, err := h.movieUsecase.ListMovie(ctx, req.GetPageSize(), req.GetPageToken())
	if err != nil {
		return nil, err
	}

	moviepbs := make([]*v1.Movie, len(movies))
	for i, movie := range movies {
		moviepbs[i] = convertMovieToPB(movie)
	}

	return &v1.ListMovieResponse{
		Movies:        moviepbs,
		NextPageToken: nextPageToken,
	}, nil
}

func (h *Handler) CreateMovie(ctx context.Context, req *v1.CreateMovieRequest) (*v1.CreateMovieResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.CreateMovie")
	defer span.End()

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(http.StatusText(http.StatusBadRequest), err.Error())
	}

	id, err := h.movieUsecase.CreateMovie(ctx, domain.CreateMovieParam{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Rating:      req.GetRating(),
		Image:       req.GetImage(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateMovieResponse{
		Id: id,
	}, nil
}

func (h *Handler) GetMovie(ctx context.Context, req *v1.GetMovieRequest) (*v1.Movie, error) {
	ctx, span := tracer.Start(ctx, "handler.GetMovie")
	defer span.End()

	movie, err := h.movieUsecase.GetMovie(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return convertMovieToPB(movie), nil
}

func (h *Handler) UpdateMovie(ctx context.Context, req *v1.UpdateMovieRequest) (*v1.UpdateMovieResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.UpdateMovie")
	defer span.End()

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(http.StatusText(http.StatusBadRequest), err.Error())
	}

	err := h.movieUsecase.UpdateMovie(ctx, domain.UpdateMovieParam{
		ID:          req.GetId(),
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMovieResponse{
		Message: fmt.Sprintf("Movie %d is updated", req.GetId()),
	}, nil
}

func (h *Handler) DeleteMovie(ctx context.Context, req *v1.DeleteMovieRequest) (*v1.DeleteMovieResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.DeleteMovie")
	defer span.End()

	if err := h.movieUsecase.DeleteMovie(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &v1.DeleteMovieResponse{
		Message: fmt.Sprintf("Movie %d is deleted", req.GetId()),
	}, nil
}

func convertMovieToPB(movie domain.Movie) *v1.Movie {
	return &v1.Movie{
		Id:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		CreatedAt:   movie.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   movie.UpdatedAt.Format(time.RFC3339),
	}
}
