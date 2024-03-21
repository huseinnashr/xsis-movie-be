package account

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/pagination"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/querybuilder"
	"github.com/huseinnashr/xsis-movie-be/internal/pkg/tracer"
)

func (r *Repo) ListMovie(ctx context.Context, pageSize int32, pageToken string) ([]domain.Movie, string, error) {
	ctx, span := tracer.Start(ctx, "repo.ListMovie")
	defer span.End()

	qb := querybuilder.New(
		"SELECT id, title, description, rating, image, created_at, updated_at FROM movies WHERE 1=1",
	)

	var cursor *ListMovieCursor
	if pageToken != "" {
		if err := pagination.DecodeCursor(pageToken, &cursor); err != nil {
			return nil, "", errors.BadRequest("cursor can't be decoded", "invalid cursor")
		}
	}

	if cursor != nil {
		qb.AddQuery("AND id <= ?", cursor.ID)
	}

	qb.AddQuery("ORDER BY id DESC LIMIT ?", pageSize+1)

	rows, err := r.sqlDatabase.QueryContext(ctx, qb.GetQuery(), qb.GetArgs()...)
	if err != nil {
		return nil, "", err
	}

	movies := []domain.Movie{}
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(
			&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image,
			&movie.CreatedAt, &movie.UpdatedAt,
		)
		if err != nil {
			return nil, "", err
		}

		movies = append(movies, movie)
	}

	nextPageToken := ""
	if len(movies) >= int(pageSize)+1 {
		nextPageToken = pagination.EncodeCursor(ListMovieCursor{
			ID: movies[len(movies)-1].ID,
		})
		movies = movies[:pageSize]
	}

	return movies, nextPageToken, nil
}

func (r *Repo) CreateMovie(ctx context.Context, param domain.CreateMovieParam) (int64, error) {
	ctx, span := tracer.Start(ctx, "repo.CreateMovie")
	defer span.End()

	row := r.sqlDatabase.QueryRowContext(ctx,
		"INSERT INTO movies(title, description, rating, image) VALUES ($1, $2, $3, $4) RETURNING id",
		param.Title, param.Description, param.Rating, param.Image,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetMovie(ctx context.Context, id int64) (domain.Movie, error) {
	ctx, span := tracer.Start(ctx, "repo.GetMovie")
	defer span.End()

	row := r.sqlDatabase.QueryRowContext(ctx,
		"SELECT id, title, description, rating, image, created_at, updated_at FROM movies WHERE id = $1",
		id,
	)

	if err := row.Err(); err != nil {
		return domain.Movie{}, err
	}

	movie := domain.Movie{}

	err := row.Scan(
		&movie.ID, &movie.Title, &movie.Description, &movie.Rating, &movie.Image,
		&movie.CreatedAt, &movie.UpdatedAt,
	)
	if err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}

func (r *Repo) UpdateMovie(ctx context.Context, param domain.UpdateMovieParam) error {
	ctx, span := tracer.Start(ctx, "repo.UpdateMovie")
	defer span.End()

	_, err := r.sqlDatabase.ExecContext(ctx, `
		UPDATE movies SET
			title = COALESCE($1, title),
			description = COALESCE($2, description),
			rating = COALESCE($3, rating),
			image = COALESCE($4, image)
		WHERE id = $5
	`, param.Title, param.Description, param.Rating, param.Image, param.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteMovie(ctx context.Context, id int64) error {
	ctx, span := tracer.Start(ctx, "repo.DeleteMovie")
	defer span.End()

	_, err := r.sqlDatabase.ExecContext(ctx,
		"DELETE FROM movies WHERE id = $1", id,
	)
	if err != nil {
		return err
	}

	return nil
}
