package account

import (
	"testing"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type UsecaseMock struct {
	movieRepo *domain.MockIMovieRepo
}

func NewMock(ctrl *gomock.Controller) *UsecaseMock {
	return &UsecaseMock{
		movieRepo: domain.NewMockIMovieRepo(ctrl),
	}
}

func (rm *UsecaseMock) toUsecase() *Usecase {
	return &Usecase{
		movieRepo: rm.movieRepo,
	}
}

func Test_NewMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	usecaseMock := NewMock(ctrl)

	usecase := usecaseMock.toUsecase()

	assert.NotNil(t, usecaseMock.movieRepo)
	assert.NotNil(t, usecase.movieRepo)
	assert.Equal(t, usecase, New(usecase.movieRepo))
}
