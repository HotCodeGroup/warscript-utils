package middlewares

import (
	"net/http"
	"strconv"
)

// чтобы иметь доступ к коду ответа
type statusWriter struct {
	http.ResponseWriter
	status string
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = strconv.Itoa(status)
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == "" {
		w.status = "200"
	}

	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}
