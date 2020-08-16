package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"testTask/internal/pkg/proto"
)

func main()  {
	grcpConn, err := grpc.Dial(
		":8001",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to session grpc")
	}
	defer grcpConn.Close()

	client := proto.NewSubServiceClient(grcpConn)

	ctx := context.Background()
	//defer cancel()
	stream, err := client.SubscribeOnSportsLines(ctx)
	if err != nil {
		log.Println(err)
	}

	err = stream.Send(&proto.SubRequest{Sports:[]string{"baseball"}, Seconds: 2})
	if err != nil {
		fmt.Println("cannot receive stream response: %v", err)
	}

	for i := 0; i < 5; i++ {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more responses")
			return
		}
		log.Print("received response: ", res.Lines)
	}

	err = stream.Send(&proto.SubRequest{Sports:[]string{"soccer", "footbal"}, Seconds: 1})
	if err != nil {
		fmt.Println("cannot receive stream response: %v", err)
	}
	for i := 0; i < 5; i++ {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more responses")
			return
		}
		log.Print("received response: ", res.Lines)
	}

	err = stream.Send(&proto.SubRequest{Sports:[]string{"soccer", "footbal", "baseball"}, Seconds: 1})
	if err != nil {
		fmt.Println("cannot receive stream response: %v", err)
	}

	for i := 0; i < 5; i++ {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more responses")
			return
		}
		log.Print("received response: ", res.Lines)
	}
}
