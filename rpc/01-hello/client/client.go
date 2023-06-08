package main

import (
	"log"
	"net/rpc"
)

func main() {
	// 连接服务
	conn, err := rpc.Dial("tcp", "127.0.0.1:9900")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	// 调用远程函数
	err = conn.Call("HelloService.Hello", "GGG", &reply)
	if err != nil {
		log.Fatal(err)
	}
	// 打印调用结果
	log.Println("call success, reply: ", reply)
}
