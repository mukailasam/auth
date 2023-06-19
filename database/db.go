package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func PgDBConnection(host, port, user, password, dbName, sslmode string) *sql.DB {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, password, sslmode)

	dbConn, err := sql.Open("pgx", dataSource)
	if err != nil {
		log.Panic(err)
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetConnMaxLifetime(1 * time.Second)
	dbConn.SetConnMaxIdleTime(30 * time.Second)

	return dbConn

}
