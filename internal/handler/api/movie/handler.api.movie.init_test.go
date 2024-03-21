package account

import (
	"testing"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type HandlerMock struct {
	movieUsecase *domain.MockIMovieUsecase
}

func NewMock(ctrl *gomock.Controller) *HandlerMock {
	return &HandlerMock{
		movieUsecase: domain.NewMockIMovieUsecase(ctrl),
	}
}

func (rm *HandlerMock) toHandler() *Handler {
	return &Handler{
		movieUsecase: rm.movieUsecase,
	}
}

func Test_NewMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	handlerMock := NewMock(ctrl)
	handler := handlerMock.toHandler()

	assert.NotNil(t, handlerMock.movieUsecase)
	assert.NotNil(t, handler.movieUsecase)
	assert.Equal(t, handler, New(handler.movieUsecase))
}
