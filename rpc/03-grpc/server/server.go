package main

import (
	"context"
	"learn-rpc/03-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello, " + req.Name}, nil
}
func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":9900")
	if err != nil {
		panic("failed listen: " + err.Error())
	}
	log.Fatal(g.Serve(lis))
}
