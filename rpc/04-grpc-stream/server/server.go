package main

import (
	"fmt"
	"learn-rpc/04-grpc-stream/proto"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) mustEmbedUnimplementedGreeterServer() {}
func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	for i := 0; i < 3; i++ {
		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("get server: %s", time.Now().String()),
		})
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *Server) PutStream(put proto.Greeter_PutStreamServer) error {
	for {
		msg, err := put.Recv()
		if err != nil {
			log.Println("err: ", err)
			break
		}
		log.Println("mgs:", msg)
	}
	return nil
}

func (s *Server) AllStream(all proto.Greeter_AllStreamServer) error {
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
			err := all.Send(&proto.StreamResData{
				Data: fmt.Sprintf("all server: %s", time.Now().String()),
			})
			if err != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatal("list")
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	log.Fatal(s.Serve(lis))

}
