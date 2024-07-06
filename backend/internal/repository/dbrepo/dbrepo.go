package dbrepo

import (
	"database/sql"
	"time"
)

type PostgresDbRepo struct {
	Conn *sql.DB
}

const dbTimeout = time.Second * 3

func (ps *PostgresDbRepo) Connection() *sql.DB {
	return ps.Conn
}
