package account

import (
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
)

type Usecase struct {
	movieRepo domain.IMovieRepo
}

func New(movieRepo domain.IMovieRepo) domain.IMovieUsecase {
	return &Usecase{
		movieRepo: movieRepo,
	}
}
