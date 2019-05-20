package logging

import (
	"io"

	"github.com/jcftang/logentriesrus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewLogger создание новго объекта логгера для нужной функции
func NewLogger(out io.Writer, token string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(out)

	// собираем логи в хранилище
	le, err := logentriesrus.NewLogentriesrusHook(token)
	if err != nil {
		return nil, errors.Wrap(err, "can not create logrus logger")
	}

	logger.AddHook(le)

	return logger, nil
}
