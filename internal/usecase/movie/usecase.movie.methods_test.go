package account

import (
	"context"
	"testing"
	"time"

	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ListMovie(t *testing.T) {
	type Arg struct {
		pageSize  int32
		pageToken string
	}

	type Test struct {
		name    string
		args    Arg
		mock    func(rm *UsecaseMock)
		want1   []domain.Movie
		want2   string
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)
	movies := []domain.Movie{
		{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock},
	}

	tests := []Test{
		{
			name: "Success",
			args: Arg{
				pageSize:  1,
				pageToken: "A",
			},
			mock: func(rm *UsecaseMock) {
				rm.movieRepo.EXPECT().ListMovie(gomock.Any(), int32(1), "A").Return(movies, "A", nil)
			},
			wantErr: false,
			want1:   movies,
			want2:   "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toUsecase()

			tt.mock(usecaseMock)
			act1, act2, err := usecase.ListMovie(context.Background(), tt.args.pageSize, tt.args.pageToken)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want1, act1)
			assert.Equal(t, tt.want2, act2)
		})
	}

}

func Test_CreateMovie(t *testing.T) {
	type Test struct {
		name    string
		args    domain.CreateMovieParam
		mock    func(rm *UsecaseMock)
		want1   int64
		wantErr bool
	}

	tests := []Test{
		{
			name: "Success",
			args: domain.CreateMovieParam{Title: "A"},
			mock: func(rm *UsecaseMock) {
				rm.movieRepo.EXPECT().CreateMovie(gomock.Any(), domain.CreateMovieParam{Title: "A"}).Return(int64(1), nil)
			},
			want1:   1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toUsecase()

			tt.mock(usecaseMock)
			act1, err := usecase.CreateMovie(context.Background(), tt.args)
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
		args    int64
		mock    func(rm *UsecaseMock)
		want1   domain.Movie
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)
	movie := domain.Movie{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock}

	tests := []Test{
		{
			name: "Success",
			args: 1,
			mock: func(rm *UsecaseMock) {
				rm.movieRepo.EXPECT().GetMovie(gomock.Any(), int64(1)).Return(movie, nil)
			},
			want1:   movie,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toUsecase()

			tt.mock(usecaseMock)
			act1, err := usecase.GetMovie(context.Background(), tt.args)
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
		args    domain.UpdateMovieParam
		mock    func(rm *UsecaseMock)
		wantErr bool
	}
	tests := []Test{
		{
			name: "Success",
			args: domain.UpdateMovieParam{ID: 1},
			mock: func(rm *UsecaseMock) {
				rm.movieRepo.EXPECT().UpdateMovie(gomock.Any(), domain.UpdateMovieParam{ID: 1}).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toUsecase()

			tt.mock(usecaseMock)
			err := usecase.UpdateMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
		})
	}
}

func Test_DeleteMovie(t *testing.T) {
	type Test struct {
		name    string
		args    int64
		mock    func(rm *UsecaseMock)
		wantErr bool
	}
	tests := []Test{
		{
			name: "Success",
			args: 1,
			mock: func(rm *UsecaseMock) {
				rm.movieRepo.EXPECT().DeleteMovie(gomock.Any(), int64(1)).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			usecaseMock := NewMock(ctrl)
			usecase := usecaseMock.toUsecase()

			tt.mock(usecaseMock)
			err := usecase.DeleteMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
		})
	}
}
