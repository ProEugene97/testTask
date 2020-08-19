package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/models"
)

type Handler struct {
	db      database.IDatabase
	logger  *zap.Logger
	isReady bool
	counter *int
	workers int
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
				zap.Any("reqId", r.Context().Value(models.ContextKey{})),
			)

			w.WriteHeader(http.StatusAccepted)
			err := json.NewEncoder(w).Encode(map[string]string{"status": "storages are not synchronized"})
			if err != nil {
				h.logger.Error(
					err.Error(),
					zap.String("func", "Status()"),
					zap.Any("reqId", r.Context().Value(models.ContextKey{})),
				)
			}
			return
		}
		h.isReady = true
	}

	err := h.db.Ping()
	if err != nil {
		h.logger.Error(err.Error(),
			zap.Any("reqId", r.Context().Value(models.ContextKey{})),
		)

		w.WriteHeader(http.StatusAccepted)
		err = json.NewEncoder(w).Encode(map[string]string{"status": "database is not available"})
		if err != nil {
			h.logger.Error(
				err.Error(),
				zap.String("func", "Status()"),
				zap.Any("reqId", r.Context().Value(models.ContextKey{})),
			)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	if err != nil {
		h.logger.Error(
			err.Error(),
			zap.String("func", "Status()"),
			zap.Any("reqId", r.Context().Value(models.ContextKey{})),
		)
	}
}
