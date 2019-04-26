package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/HotCodeGroup/warscript-utils/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// RecoverMiddleware ловит паники и кидает 500ки
func RecoverMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithField("method", "RECOVER").Error(err)
				utils.NewErrorResponseWriter(w, log.WithField("method", "RecoverMiddleware")).
					WriteError(http.StatusInternalServerError, utils.ErrInternal)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// AccessLogMiddleware логирование всех запросов
func AccessLogMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := uuid.NewV4()
		ctx := context.WithValue(r.Context(), utils.RequestUUIDKey, token.String()[:8])

		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		log.WithFields(logrus.Fields{
			"token":     token.String()[:8],
			"method":    r.Method,
			"work_time": time.Since(start).Seconds(),
		}).Info(r.URL.Path)
	})
}
