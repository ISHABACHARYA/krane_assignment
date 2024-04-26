package postgres

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func New(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GoquNew(db *sql.DB) *goqu.Database {
	return goqu.New("postgres", db)
}
