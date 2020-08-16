package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testTask/internal/pkg/models"

	//"strings"
	"testTask/internal/pkg/database"
	//"testTask/internal/pkg/models"
	"time"
)

type Worker struct {
	timeOut   int
	sport     string
	db 	      database.IDatabase
	isFirst   bool
	readyChan chan interface{}
}

func NewWorker(timeOut int, sport string, db database.IDatabase, readyChan chan interface{}) *Worker {
	return &Worker{
		timeOut,
		sport,
		db,
		false,
		readyChan,
	}
}

func (w *Worker) Run() {
	for {
		resp, err := http.Get("http://localhost:8000/api/v1/lines/" + w.sport)
		if err != nil {
			fmt.Println(err)
		}

		m := map[string]map[string]string {}
		err = json.NewDecoder(resp.Body).Decode(&m)
		if err != nil {
			fmt.Println(err)
			return
		}

		lines, ok := m["lines"]
		if !ok {
			continue
		}

		coef, ok := lines[strings.ToUpper(w.sport)]
		if !ok {
			continue
		}

		err = w.db.Set(models.Line{ Sport: w.sport, Coef: coef})
		if err != nil {
			fmt.Println(err)
		}

		if !w.isFirst {
			w.isFirst = true
			w.readyChan <- true
			w.readyChan <- true
		}

		timer := time.NewTimer(time.Duration(w.timeOut) * time.Second)
		<-timer.C
	}
}
