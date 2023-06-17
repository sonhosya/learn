package main

import (
	"context"
	"log"
	"net"

	"learn-rpc/05-grpc-metadata/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HelloServer struct {
	proto.UnimplementedHelloServer
}

func (s HelloServer) SayHello(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	// 增加处理 metadata
	md, ok := metadata.FromIncomingContext(ctx)
	log.Printf(" metadata ok: %v\n", ok)
	for k, v := range md {
		log.Printf("metadata k,v: %v=%v\n", k, v)
	}
	log.Println("recv: ", req)
	return &proto.Resp{Msg: "hello, " + req.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterHelloServer(g, &HelloServer{})
	lis, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatalf("listen failed, err: %s", err)
	}
	log.Fatalf("grpc serve failed,err: %s", g.Serve(lis))
}
