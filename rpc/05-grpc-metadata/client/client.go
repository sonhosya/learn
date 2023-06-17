package main

import (
	"context"
	"log"
	"time"

	"learn-rpc/05-grpc-metadata/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9900", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial failed, err: %s", err)
	}
	defer conn.Close()
	cli := proto.NewHelloClient(conn)
	// 增加 metadata
	md := metadata.New(map[string]string{"name": "GGG", "timestamp": time.Now().Local().String()})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := cli.SayHello(ctx, &proto.Req{Name: "GGG"})
	if err != nil {
		log.Printf("call SayHello Error: %s\n", err)
	}
	log.Println("resp: ", resp)
}
