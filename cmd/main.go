package main

import (
	"go.uber.org/zap"
	"log"
	"sync"
	"testTask/internal/app/grpc"
	"testTask/internal/app/http"
	"testTask/internal/pkg/config"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/logger"
	"testTask/internal/pkg/worker"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal("reading config error")
	}

	rdb := database.NewRedisDB(database.NewPool(c.Redis))
	l := logger.NewLogger(c.LogLevel)

	mu := &sync.Mutex{}
	counter := 0
	for i, sport := range c.Sports {
		go worker.NewWorker(c.Provider, c.Timeouts[i], sport, rdb, l, &counter, mu).Run()
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := http.NewServer(c.HTTPHost+":"+c.HTTPPort, rdb, l, &counter, len(c.Sports)).Start()
		if err != nil {
			l.Error(err.Error(),
				zap.String("func", "HTTPServer"),
			)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			if counter == len(c.Sports) {
				break
			}
		}
		err := grpc.NewServer(c.GRPCHost+":"+c.GRPCPort, rdb, l).Start()
		if err != nil {
			l.Error(err.Error(),
				zap.String("func", "GRPCServer"),
			)
		}
	}(wg)

	wg.Wait()
}
