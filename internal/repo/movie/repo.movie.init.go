package account

import (
	"github.com/huseinnashr/xsis-movie-be/internal/config"
	"github.com/huseinnashr/xsis-movie-be/internal/domain"
)

type Repo struct {
	config      *config.Config
	sqlDatabase domain.ISQLDatabase
}

func New(config *config.Config, sqlDatabase domain.ISQLDatabase) domain.IMovieRepo {
	return &Repo{
		config:      config,
		sqlDatabase: sqlDatabase,
	}
}
