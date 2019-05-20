package postgresql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// Queryer интерфес для передачи в методы запроса транзаций или коннекта к базе
type Queryer interface {
	QueryRow(string, ...interface{}) *sql.Row
}

// Connect подключение к базе данных PostgreSQL
func Connect(dbUser, dbPass, dbHost, dbPort, dbName string) (*sql.DB, error) {
	var err error
	_, err = strconv.ParseInt(dbPort, 10, 16)
	if err != nil {
		return nil, errors.Wrap(err, "port int parse error")
	}

	pqConn, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return nil, errors.Wrap(err, "can not create pq connection")
	}

	return pqConn, nil
}
