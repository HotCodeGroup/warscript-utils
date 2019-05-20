package rabbitmq

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const (
	amqpPattern = "amqp://%s:%s@%s:%s/"
)

// Connect подключение к очереди RabbitMQ
func Connect(dbUser, dbPass, dbHost, dbPort string) (*amqp.Connection, error) {
	var err error
	_, err = strconv.ParseInt(dbPort, 10, 16)
	if err != nil {
		return nil, errors.Wrap(err, "port int parse error")
	}

	rabbitConn, err := amqp.Dial(fmt.Sprintf(amqpPattern, dbUser, dbPass, dbHost, dbPort))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to RabbitMQ")
	}

	return rabbitConn, nil
}
