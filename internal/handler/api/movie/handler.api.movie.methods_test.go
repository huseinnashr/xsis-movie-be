package account

import (
	"context"
	"errors"
	"testing"
	"time"

	v1 "github.com/huseinnashr/xsis-movie-be/api/v1"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ListMovie(t *testing.T) {
	type Test struct {
		name    string
		args    *v1.ListMovieRequest
		mock    func(rm *HandlerMock)
		want1   *v1.ListMovieResponse
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)
	movies := []domain.Movie{
		{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock},
	}

	tests := []Test{
		{
			name:    "Input invalid",
			args:    &v1.ListMovieRequest{PageSize: -1},
			mock:    func(rm *HandlerMock) {},
			wantErr: true,
		},
		{
			name: "List movie error",
			args: &v1.ListMovieRequest{PageSize: 1, PageToken: "A"},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().ListMovie(gomock.Any(), int32(1), "A").Return(nil, "", errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: &v1.ListMovieRequest{PageSize: 1, PageToken: "A"},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().ListMovie(gomock.Any(), int32(1), "A").Return(movies, "A", nil)
			},
			wantErr: false,
			want1: &v1.ListMovieResponse{
				Movies:        []*v1.Movie{convertMovieToPB(movies[0])},
				NextPageToken: "A",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handlerMock := NewMock(ctrl)
			handler := handlerMock.toHandler()

			tt.mock(handlerMock)
			act1, err := handler.ListMovies(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
		})
	}

}

func Test_CreateMovie(t *testing.T) {
	type Test struct {
		name    string
		args    *v1.CreateMovieRequest
		mock    func(rm *HandlerMock)
		want1   *v1.CreateMovieResponse
		wantErr bool
	}

	tests := []Test{
		{
			name:    "Validate error",
			args:    &v1.CreateMovieRequest{Rating: -1},
			mock:    func(rm *HandlerMock) {},
			wantErr: true,
		},
		{
			name: "Create movie error",
			args: &v1.CreateMovieRequest{Title: "A"},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().CreateMovie(gomock.Any(), domain.CreateMovieParam{Title: "A"}).Return(int64(0), errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: &v1.CreateMovieRequest{Title: "A"},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().CreateMovie(gomock.Any(), domain.CreateMovieParam{Title: "A"}).Return(int64(1), nil)
			},
			want1: &v1.CreateMovieResponse{
				Id: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handlerMock := NewMock(ctrl)
			handler := handlerMock.toHandler()

			tt.mock(handlerMock)
			act1, err := handler.CreateMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
		})
	}
}

func Test_GetMovie(t *testing.T) {
	type Test struct {
		name    string
		args    *v1.GetMovieRequest
		mock    func(rm *HandlerMock)
		want1   *v1.Movie
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)
	movie := domain.Movie{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock}

	tests := []Test{
		{
			name: "Get movie error",
			args: &v1.GetMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().GetMovie(gomock.Any(), int64(1)).Return(domain.Movie{}, errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: &v1.GetMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().GetMovie(gomock.Any(), int64(1)).Return(movie, nil)
			},
			want1:   convertMovieToPB(movie),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handlerMock := NewMock(ctrl)
			handler := handlerMock.toHandler()

			tt.mock(handlerMock)
			act1, err := handler.GetMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
		})
	}
}

func Test_UpdateMovie(t *testing.T) {
	type Test struct {
		name    string
		args    *v1.UpdateMovieRequest
		mock    func(rm *HandlerMock)
		want1   *v1.UpdateMovieResponse
		wantErr bool
	}

	invalidRating := float32(-1)
	tests := []Test{
		{
			name:    "Validate error",
			args:    &v1.UpdateMovieRequest{Rating: &invalidRating},
			mock:    func(rm *HandlerMock) {},
			wantErr: true,
		},
		{
			name: "Update movie error",
			args: &v1.UpdateMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().UpdateMovie(gomock.Any(), domain.UpdateMovieParam{ID: 1}).Return(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: &v1.UpdateMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().UpdateMovie(gomock.Any(), domain.UpdateMovieParam{ID: 1}).Return(nil)
			},
			want1: &v1.UpdateMovieResponse{
				Message: "Movie 1 is updated",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toHandler()

			tt.mock(usecaseMock)
			act1, err := usecase.UpdateMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
		})
	}
}

func Test_DeleteMovie(t *testing.T) {
	type Test struct {
		name    string
		args    *v1.DeleteMovieRequest
		mock    func(rm *HandlerMock)
		want1   *v1.DeleteMovieResponse
		wantErr bool
	}
	tests := []Test{
		{
			name: "Delete movie error",
			args: &v1.DeleteMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().DeleteMovie(gomock.Any(), int64(1)).Return(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: &v1.DeleteMovieRequest{Id: 1},
			mock: func(rm *HandlerMock) {
				rm.movieUsecase.EXPECT().DeleteMovie(gomock.Any(), int64(1)).Return(nil)
			},
			want1: &v1.DeleteMovieResponse{
				Message: "Movie 1 is deleted",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handlerMock := NewMock(ctrl)
			handler := handlerMock.toHandler()

			tt.mock(handlerMock)
			act1, err := handler.DeleteMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
		})
	}
}
