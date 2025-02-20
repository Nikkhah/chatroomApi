package repository

import (
	"chatroom/internal/core/port/repository"
	"chatroom/internal/infra/config"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	*sql.DB
}

func NewDB(conf config.DatabaseConfig) (repository.Database, error) {
	db, err := newDatabase(conf)
	if err != nil {
		return nil, err
	}
	return &database{
		db,
	}, nil
}

func newDatabase(conf config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(conf.Driver, conf.Url)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(conf.ConnMaxLifetimeInMinute))
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func (da database) GetDB() *sql.DB {
	return da.DB
}
