package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"testTask/internal/pkg/api"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/proto"
	"testTask/internal/pkg/worker"
)

func main() {
	rdb := database.NewRedisDB(database.NewPool("localhost:6379"))

	sports := []string{"baseball", "soccer", "football"}
	m :=  map[chan interface{}]bool{}
	for _, sport := range sports {
		ch := make(chan interface{}, 2)
		m[ch] = false
		go worker.NewWorker(3, sport, rdb, ch).Run()
	}

	handler := api.NewHandler(rdb, m)
	http.HandleFunc("/ready", handler.Status)
	go http.ListenAndServe(":8080", nil)

	server := grpc.NewServer()

	proto.RegisterSubServiceServer(server, api.NewSubService(rdb))

	for ch := range m {
		<- ch
	}
	fmt.Println("here")

	lis, ok := net.Listen("tcp", ":8001")
	if ok != nil {
		log.Fatalln("cant listet port", ok)
	}
	server.Serve(lis)

}
