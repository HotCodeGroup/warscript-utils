package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/HotCodeGroup/warscript-utils/metrics"

	"github.com/HotCodeGroup/warscript-utils/models"
	"github.com/HotCodeGroup/warscript-utils/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

const (
	// SessionInfoKey ключ, по которому в контексте
	// реквеста хранится структура юзера после валидации
	SessionInfoKey utils.ContextKey = 1
)

var hits = metrics.NewHits()

// WithAuthentication проверка токена перед исполнением запроса
//nolint: interfacer
func WithAuthentication(next http.HandlerFunc, l *logrus.Logger, cli models.AuthClient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := utils.GetLogger(r, l, "CheckUsername")
		errWriter := utils.NewErrorResponseWriter(w, logger)

		cookie, err := r.Cookie("JSESSIONID")
		if err != nil || cookie == nil {
			errWriter.WriteWarn(http.StatusUnauthorized, errors.Wrap(err, "can not load cookie"))
			return
		}

		session, err := cli.GetSessionInfo(r.Context(), &models.SessionToken{Token: cookie.Value})
		if err != nil {
			errWriter.WriteError(http.StatusUnauthorized, errors.Wrap(err, "get session error"))
			return

		}

		ctx := context.WithValue(r.Context(), SessionInfoKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
		logger := utils.GetLogger(r, log, "AccessLogMiddleware")

		token := uuid.NewV4()
		ctx := context.WithValue(r.Context(), utils.RequestUUIDKey, token.String()[:8])

		start := time.Now()
		uw, ok := w.(responseWriter)
		if !ok {
			logger.Error("response does not implement http.Hijacker")
			utils.NewErrorResponseWriter(w, logger).
				WriteError(http.StatusInternalServerError, utils.ErrInternal)
		}

		sw := statusWriter{responseWriter: uw}

		next.ServeHTTP(&sw, r.WithContext(ctx))
		hits.WithLabelValues(sw.status, r.URL.String()).Inc()
		log.WithFields(logrus.Fields{
			"token":     token.String()[:8],
			"status":    sw.status,
			"method":    r.Method,
			"work_time": time.Since(start).Seconds(),
		}).Info(r.URL.Path)
	})
}

// WithLimiter для запросов, у которых есть ограничение в секунду
//nolint: interfacer
func WithLimiter(next http.HandlerFunc, limiter *rate.Limiter, log *logrus.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			log.WithField("method", "WithLimiter").Warn("too many requests")
			http.Error(w, "", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
