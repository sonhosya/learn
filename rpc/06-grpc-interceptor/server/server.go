package main

import (
	"context"
	"log"
	"net"

	"learn-rpc/06-grpc-interceptor/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	// 增加 服务端 拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("接收到一个新请求")
		var user, passwd string
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "缺少认证信息")
		}
		if v, ok := md["user"]; ok {
			user = v[0]
		}

		if v, ok := md["passwd"]; ok {
			passwd = v[0]
		}
		if user != "GGG" || passwd != "ggg" {
			return resp, status.Error(codes.Unauthenticated, "认证不通过")
		}
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterHelloServer(g, &HelloServer{})
	lis, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatalf("listen failed, err: %s", err)
	}
	log.Fatalf("grpc serve failed,err: %s", g.Serve(lis))
}
