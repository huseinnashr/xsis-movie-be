package domain

import (
	"context"
	"database/sql"
	"time"
)

//go:generate mockgen -source=./domain.interface.go -destination=./domain.interface.mock.go -package=domain
type IMovieUsecase interface {
	ListMovie(ctx context.Context, size int32, token string) ([]Movie, string, error)
	CreateMovie(ctx context.Context, param CreateMovieParam) (int64, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
	UpdateMovie(ctx context.Context, param UpdateMovieParam) error
	DeleteMovie(ctx context.Context, id int64) error
}

type IMovieRepo interface {
	ListMovie(ctx context.Context, size int32, token string) ([]Movie, string, error)
	CreateMovie(ctx context.Context, param CreateMovieParam) (int64, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
	UpdateMovie(ctx context.Context, param UpdateMovieParam) error
	DeleteMovie(ctx context.Context, id int64) error
}

type ISQLDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Movie struct {
	ID          int64
	Title       string
	Description string
	Rating      float32
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateMovieParam struct {
	Title       string
	Description string
	Rating      float32
	Image       string
}

type UpdateMovieParam struct {
	ID          int64
	Title       *string
	Description *string
	Rating      *float32
	Image       *string
}
