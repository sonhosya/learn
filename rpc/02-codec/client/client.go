package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 连接服务
	conn, err := net.Dial("tcp", "127.0.0.1:9900")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var reply string

	// 调用远程函数
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	defer client.Close()

	err = client.Call("HelloService.Hello", "GGG", &reply)
	if err != nil {
		log.Fatal(err)
	}
	// 打印调用结果
	log.Println("call success, reply: ", reply)
}
