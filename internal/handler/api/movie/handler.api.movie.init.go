package account

import (
	v1 "github.com/huseinnashr/xsis-movie-be/api/v1"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
)

type Handler struct {
	v1.UnimplementedAccountServiceServer
	movieUsecase domain.IMovieUsecase
}

func New(movieUsecase domain.IMovieUsecase) v1.AccountServiceServer {
	return &Handler{
		movieUsecase: movieUsecase,
	}
}
