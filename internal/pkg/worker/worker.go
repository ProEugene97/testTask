package worker

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"sync"
	"testTask/internal/pkg/models"
	"testTask/internal/pkg/database"
	"time"
)

type Worker struct {
	url     string
	timeOut int
	sport   string
	db      database.IDatabase
	logger  *zap.Logger
	counter *int
	mu      *sync.Mutex
	isReady bool
}

func NewWorker(
	url string,
	timeOut int,
	sport string,
	db database.IDatabase,
	logger *zap.Logger,
	counter *int,
	mu *sync.Mutex) *Worker {

	return &Worker{
		url,
		timeOut,
		sport,
		db,
		logger,
		counter,
		mu,
		false,
	}
}

func (w *Worker) Run() {
	timeout := 0
	for {
		time.Sleep(time.Duration(timeout) * time.Second)

		timeout = w.timeOut

		resp, err := http.Get(w.url + w.sport)
		if err != nil {
			w.logger.Warn(
				err.Error(),
				zap.String("func", "http.Get()"),
				zap.String("addr", w.url+w.sport),
				zap.String("sport", w.sport),
			)
			continue
		}

		line, err := w.readResponse(resp.Body)
		if err != nil {
			w.logger.Error(
				err.Error(),
				zap.String("func", "readResponse"),
				zap.String("addr", w.url+w.sport),
				zap.String("sport", w.sport),
			)
			continue
		}

		err = w.db.Set(line)
		if err != nil {
			w.logger.Error(
				err.Error(),
				zap.String("func", "db.set"),
				zap.String("addr", w.url+w.sport),
				zap.String("sport", w.sport),
			)
			continue
		}

		if !w.isReady {
			w.isReady = true
			w.mu.Lock()
			*w.counter++
			w.mu.Unlock()

		}
	}
}

func (w *Worker) readResponse(body io.ReadCloser) (*models.Line, error) {
	m := map[string]map[string]string{}
	err := json.NewDecoder(body).Decode(&m)
	if err != nil {
		return nil, errors.Wrap(err, "Decoding error: ")
	}

	lines, ok := m["lines"]
	if !ok {
		return nil, errors.Wrap(err, "There isn't lines field: ")
	}

	coef, ok := lines[strings.ToUpper(w.sport)]
	if !ok {
		return nil, errors.Wrap(err, "There isn't sport field: ")
	}

	return &models.Line{Sport: w.sport, Coef: coef}, nil
}
