package account

import (
	"database/sql"

	"github.com/huseinnashr/xsis-movie-be/internal/config"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
)

type Repo struct {
	config      *config.Config
	sqlDatabase *sql.DB
}

func New(config *config.Config, sqlDatabase *sql.DB) domain.IMovieRepo {
	return &Repo{
		config:      config,
		sqlDatabase: sqlDatabase,
	}
}
