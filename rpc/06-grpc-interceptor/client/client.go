package main

import (
	"context"
	"log"
	"time"

	"learn-rpc/06-grpc-interceptor/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 增加 客服端 拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("本次执行耗时： %s\n", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("127.0.0.1:9900", grpc.WithInsecure(), opt)
	if err != nil {
		log.Fatalf("grpc dial failed, err: %s", err)
	}
	defer conn.Close()
	cli := proto.NewHelloClient(conn)
	resp, err := cli.SayHello(context.Background(), &proto.Req{Name: "GGG"})
	if err != nil {
		log.Printf("call SayHello Error: %s\n", err)
	}
	log.Println("resp: ", resp)

	// 增加 metadata
	md := metadata.New(map[string]string{"user": "GGG", "passwd": "ggg", "timestamp": time.Now().Local().String()})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err = cli.SayHello(ctx, &proto.Req{Name: "GGG"})
	if err != nil {
		log.Printf("call SayHello Error: %s\n", err)
	}
	log.Println("resp: ", resp)
}
