package models

import (
	"database/sql"

	"github.com/boj/redistore"
)

type Model struct {
	PgDBConn  *sql.DB
	RediStore *redistore.RediStore
}
