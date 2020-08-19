package http

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"testTask/internal/pkg/api"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/middleware"
	"time"
)

type HTTPServer struct {
	addr    string
	db      database.IDatabase
	logger  *zap.Logger
	counter *int
	workers int
}

func NewHTTPServer(addr string, db database.IDatabase, logger *zap.Logger, counter *int, workers int) *HTTPServer {
	return &HTTPServer{
		addr,
		db,
		logger,
		counter,
		workers,
	}
}

func (hs *HTTPServer) Start() error {
	r := mux.NewRouter()
	m := middleware.NewMiddleware(hs.logger)
	handler := api.NewHandler(hs.db, hs.logger, hs.counter, hs.workers)

	r.HandleFunc("/ready", handler.Status)
	r.Use(m.PanicMiddleware)
	r.Use(m.LogMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         hs.addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	hs.logger.Info("Start http server",
		zap.String("addr", hs.addr),
		zap.Time("start", time.Now()),
	)

	return srv.ListenAndServe()
}
