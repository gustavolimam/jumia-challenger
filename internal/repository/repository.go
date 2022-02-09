package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	CustomerQueries
}

type repository struct {
	db *sql.DB
}

func New() Repository {
	db, err := sql.Open("sqlite3", "database/sample.db")
	if err != nil {
		log.Panic().Err(err).Msg("Error trying to open a database connection")
	}

	return &repository{
		db: db,
	}
}
