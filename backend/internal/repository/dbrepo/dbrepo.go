package dbrepo

import (
	"database/sql"
	"time"
)

type PostgresDbRepo struct {
	Conn *sql.DB
}

const dbTimeout = time.Second * 3

func (db *PostgresDbRepo) Connection() *sql.DB {
	return db.Conn
}
