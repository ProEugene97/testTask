package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"testTask/internal/pkg/models"
	"time"
)

type Middleware struct {
	logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) Middleware {
	return Middleware{
		logger: logger,
	}
}

func (m *Middleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := r.Context()
		reqID := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx = context.WithValue(ctx, models.ContextKey{}, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
		m.logger.Info(r.URL.Path,
			zap.String("reqId:", reqID),
			zap.String("method", r.Method),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("url", r.URL.Path),
			zap.Time("start", start),
			zap.Duration("work_time", time.Since(start)),
		)
	})
}

func (m *Middleware) PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
