package middlewares

import (
	"bufio"
	"net"
	"net/http"
	"strconv"
)

type responseWriter interface {
	http.ResponseWriter
	http.Hijacker
}

// чтобы иметь доступ к коду ответа
type statusWriter struct {
	responseWriter
	status string
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = strconv.Itoa(status)
	w.responseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == "" {
		w.status = "200"
	}

	n, err := w.responseWriter.Write(b)
	w.length += n
	return n, err
}

func (w *statusWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.responseWriter.Hijack()
}
