package postgresql

import (
	"github.com/pkg/errors"
	"strconv"

	"github.com/jackc/pgx"
)

type Queryer interface {
	QueryRow(string, ...interface{}) *pgx.Row
}

func Connect(dbUser, dbPass, dbHost, dbPort, dbName string) (*pgx.ConnPool, error) {
	var err error
	port, err := strconv.ParseInt(dbPort, 10, 16)
	if err != nil {
		return nil, errors.Wrap(err, "port int parse error")
	}

	pgxConn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     dbHost,
			User:     dbUser,
			Password: dbPass,
			Database: dbName,
			Port:     uint16(port),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "can not create pgx connection pool")
	}

	return pgxConn, nil
}
