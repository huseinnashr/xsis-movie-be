package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huseinnashr/xsis-movie-be/internal/config"
	accounthandler "github.com/huseinnashr/xsis-movie-be/internal/handler/api/movie"
	movierepo "github.com/huseinnashr/xsis-movie-be/internal/repo/movie"
	movieusecase "github.com/huseinnashr/xsis-movie-be/internal/usecase/movie"
	_ "github.com/lib/pq"
)

func startApp(ctx context.Context, config *config.Config) error {
	sqlDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Resource.SQLDatabase.Host, config.Resource.SQLDatabase.Port, config.Resource.SQLDatabase.User,
		config.Resource.SQLDatabase.Password, config.Resource.SQLDatabase.DBName,
	)
	sqlDatabase, err := sql.Open("postgres", sqlDSN)
	if err != nil {
		return err
	}

	movieRepo := movierepo.New(config, sqlDatabase)
	movieUsecase := movieusecase.New(movieRepo)
	accountHandler := accounthandler.New(movieUsecase)

	if err := startServer(ctx, config, accountHandler); err != nil {
		return err
	}

	return nil
}
