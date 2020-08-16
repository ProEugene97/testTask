package api

import (
	"net/http"
	"testTask/internal/pkg/database"
)

type Handler struct {
	db database.IDatabase
	isReady bool
	readyChans map[chan interface{}]bool
}

func NewHandler(db database.IDatabase, readyChans map[chan interface{}]bool) *Handler {
	return &Handler{
		db,
		false,
		readyChans,
	}
}


func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	if !h.isReady {
		for ch := range h.readyChans {
			if !h.readyChans[ch] {
				select {
				case <-ch:
					h.readyChans[ch] = true
				default:
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
		}
		h.isReady = true
	}

	err := h.db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
}