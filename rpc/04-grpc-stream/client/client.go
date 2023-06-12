package main

import (
	"context"
	"learn-rpc/04-grpc-stream/proto"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var client proto.GreeterClient

func main() {
	conn, err := grpc.Dial("127.0.0.1:9900", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = proto.NewGreeterClient(conn)
	res, err := client.GetStream(context.Background(), &proto.StreamReqData{Data: "get GGG"})
	if err != nil {
		log.Fatal(err)
	}
	defer res.CloseSend()
	getExec()
	putExec()
	allExec()
}
func putExec() {
	put, err := client.PutStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		err := put.Send(&proto.StreamReqData{
			Data: "put GGG"})
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	put.CloseSend()
}
func getExec() {
	res, err := client.GetStream(context.Background(), &proto.StreamReqData{Data: "get GGG"})
	if err != nil {
		log.Fatal(err)
	}
	defer res.CloseSend()
	for {
		msg, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(msg.String())
	}
	res.CloseSend()
}
func allExec() {

	all, err := client.AllStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			msg, err := all.Recv()
			if err != nil {
				log.Println("err: ", err)
				break
			}
			log.Println("mgs:", msg)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			err := all.Send(&proto.StreamReqData{
				Data: "all GGG",
			})
			if err != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		all.CloseSend()
	}()
	wg.Wait()
}
