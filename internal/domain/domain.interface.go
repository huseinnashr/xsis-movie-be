package domain

import (
	"context"
	"database/sql"
)

type IMovieUsecase interface {
	ListMovie(ctx context.Context, size int32, token string) ([]Movie, string, error)
}

type IMovieRepo interface {
	ListMovie(ctx context.Context, size int32, token string) ([]Movie, string, error)
}

type ISQLDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Movie struct {
	ID int32
}
