package api

import (
	"go.uber.org/zap"
	"net/http"
	"testTask/internal/pkg/database"
)

type Handler struct {
	db       database.IDatabase
	logger   *zap.Logger
	isReady   bool
	counter *int
	workers  int
}

func NewHandler(db database.IDatabase, logger *zap.Logger, counter *int, workers int) *Handler {
	return &Handler{
		db,
		logger,
		false,
		counter,
		workers,
	}
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	if !h.isReady {
		if *h.counter < h.workers {
			h.logger.Debug(
				"Data isn't received from provider",
				zap.String("func", "Status()"),
				zap.Any("reqId", r.Context().Value("reqId")),
				)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.isReady = true
	}

	err := h.db.Ping()
	if err != nil {
		h.logger.Error(err.Error(),
			zap.Any("reqId", r.Context().Value("reqId")),
			)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
}