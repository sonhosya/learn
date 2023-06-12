package main

import (
	"context"
	"learn-rpc/03-grpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9900", grpc.WithInsecure())
	if err != nil {
		panic("connect failed, err: " + err.Error())
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	req := proto.HelloRequest{Name: "GGG"}
	r, err := c.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal("call sayhello failed, err: ", err)
	}
	log.Println("call sayhello success, response: ", r.String())
}
