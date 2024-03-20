package account

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/huseinnashr/xsis-movie-be/api/v1"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (h *Handler) ListMovies(ctx context.Context, req *v1.ListMovieRequest) (*v1.ListMovieResponse, error) {
	_, span := tracer.Start(ctx, "handler.ListMovie")
	defer span.End()

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest(http.StatusText(http.StatusBadRequest), err.Error())
	}

	return &v1.ListMovieResponse{
		Movies: []*v1.Movie{
			{Id: 1},
		},
	}, nil
}
