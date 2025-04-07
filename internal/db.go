package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DbPsql interface {
	ConnectDb() (*sql.DB, error)
}

type DbImpl struct {}

func NewDbImpl() *DbImpl {
	return &DbImpl{}
}

func (d *DbImpl) ConnectDb() (*sql.DB, error) {
	connStr := "host=127.0.0.1 port=5433 user=root password=root dbname=messages sslmode=disable"
	db, errCon := sql.Open("postgres", connStr)
	if errCon != nil {
		return nil, errCon
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	return db, nil
}