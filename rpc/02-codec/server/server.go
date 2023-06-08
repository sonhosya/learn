package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 实例化 server
	listener, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatal(err)
	}
	// 注册处理逻辑
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		log.Fatal(err)
	}
	// 启动服务
	for {
		conn, _ := listener.Accept()
		// 使用json编解码
		// rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
