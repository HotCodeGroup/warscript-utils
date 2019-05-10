package postgresql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type Queryer interface {
	QueryRow(string, ...interface{}) *sql.Row
}

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

// func Connect(dbUser, dbPass, dbHost, dbPort, dbName string) (*pgx.ConnPool, error) {
// 	var err error
// 	port, err := strconv.ParseInt(dbPort, 10, 16)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "port int parse error")
// 	}

// 	pgxConn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
// 		ConnConfig: pgx.ConnConfig{
// 			Host:     dbHost,
// 			User:     dbUser,
// 			Password: dbPass,
// 			Database: dbName,
// 			Port:     uint16(port),
// 		},
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "can not create pgx connection pool")
// 	}

// 	return pgxConn, nil
// }
