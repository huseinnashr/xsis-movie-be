package account

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huseinnashr/xsis-movie-be/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type RepoMock struct {
	config    *config.Config
	sqlDb     *sql.DB
	sqlDbMock sqlmock.Sqlmock
}

func NewMock(ctrl *gomock.Controller) (*RepoMock, func()) {
	db, mock, _ := sqlmock.New()
	dispose := func() {
		db.Close()
	}

	return &RepoMock{
		config:    &config.Config{},
		sqlDb:     db,
		sqlDbMock: mock,
	}, dispose
}

func (rm *RepoMock) toRepo() *Repo {
	return &Repo{
		config:      rm.config,
		sqlDatabase: rm.sqlDb,
	}
}

func Test_NewMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock, dispose := NewMock(ctrl)
	defer dispose()

	repo := repoMock.toRepo()

	assert.NotNil(t, repoMock.config)
	assert.NotNil(t, repoMock.sqlDb)
	assert.NotNil(t, repoMock.sqlDbMock)
	assert.NotNil(t, repo.config)
	assert.NotNil(t, repo.sqlDatabase)
	assert.Equal(t, repo, New(repo.config, repo.sqlDatabase))
}
