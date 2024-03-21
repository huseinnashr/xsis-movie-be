package account

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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
		mock    func(rm *RepoMock)
		want1   []domain.Movie
		want2   string
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := []Test{
		{
			name: "Invalid cursor",
			args: Arg{
				pageToken: "A",
			},
			mock:    func(rm *RepoMock) {},
			wantErr: true,
		},
		{
			name: "DB query failed",
			args: Arg{
				pageToken: "eyJpZCI6MX0",
			},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnError(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Failed to scan result",
			args: Arg{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnRows(
					sqlmock.NewRows([]string{"id"}).AddRows(
						[]driver.Value{1},
					),
				)
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: Arg{
				pageSize: 1,
			},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnRows(
					sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).AddRows(
						[]driver.Value{1, "A", "B", 7, "C", timeMock, timeMock},
						[]driver.Value{2, "A", "B", 7, "C", timeMock, timeMock},
					),
				)
			},
			want1: []domain.Movie{
				{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock},
			},
			want2: "eyJpZCI6Mn0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repoMock, dispose := NewMock(ctrl)
			defer dispose()

			repo := repoMock.toRepo()

			tt.mock(repoMock)
			act1, act2, err := repo.ListMovie(context.Background(), tt.args.pageSize, tt.args.pageToken)
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
		mock    func(rm *RepoMock)
		want1   int64
		wantErr bool
	}

	tests := []Test{
		{
			name: "Query row error",
			args: domain.CreateMovieParam{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("INSERT").WillReturnError(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Scan error",
			args: domain.CreateMovieParam{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("A"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: domain.CreateMovieParam{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want1:   1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repoMock, dispose := NewMock(ctrl)
			defer dispose()

			repo := repoMock.toRepo()

			tt.mock(repoMock)
			act1, err := repo.CreateMovie(context.Background(), tt.args)
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
		mock    func(rm *RepoMock)
		want1   domain.Movie
		wantErr bool
	}

	timeMock := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	tests := []Test{
		{
			name: "Query row error",
			args: 0,
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnError(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Scan error",
			args: 0,
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("A"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: 1,
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectQuery("SELECT").WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
						AddRow(1, "A", "B", 7, "C", timeMock, timeMock),
				)
			},
			want1:   domain.Movie{ID: 1, Title: "A", Description: "B", Rating: 7, Image: "C", CreatedAt: timeMock, UpdatedAt: timeMock},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repoMock, dispose := NewMock(ctrl)
			defer dispose()

			repo := repoMock.toRepo()

			tt.mock(repoMock)
			act1, err := repo.GetMovie(context.Background(), tt.args)
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
		mock    func(rm *RepoMock)
		wantErr bool
	}
	tests := []Test{
		{
			name: "Exec query error",
			args: domain.UpdateMovieParam{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectExec("UPDATE").WillReturnError(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: domain.UpdateMovieParam{},
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repoMock, dispose := NewMock(ctrl)
			defer dispose()

			repo := repoMock.toRepo()

			tt.mock(repoMock)
			err := repo.UpdateMovie(context.Background(), tt.args)
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
		mock    func(rm *RepoMock)
		wantErr bool
	}
	tests := []Test{
		{
			name: "Exec query error",
			args: 0,
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectExec("DELETE").WillReturnError(errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: 1,
			mock: func(rm *RepoMock) {
				rm.sqlDbMock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repoMock, dispose := NewMock(ctrl)
			defer dispose()

			repo := repoMock.toRepo()

			tt.mock(repoMock)
			err := repo.DeleteMovie(context.Background(), tt.args)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
		})
	}
}
