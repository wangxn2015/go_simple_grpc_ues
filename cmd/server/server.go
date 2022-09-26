package main

import (
	"flag"
	"fmt"
	"github.com/wangxn2015/go_simple_grpc_ues/api/track_msg"
	"github.com/wangxn2015/go_simple_grpc_ues/pkg/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("hello grpc from server")
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	fmt.Printf("start server on port: %d\n", *port)

	uesServer := service.NewUesServer(service.NewInMemoryUeStore())
	grpcServer := grpc.NewServer()
	//把服务注册到grpcServer上
	track_msg.RegisterUEsServer(grpcServer, uesServer)

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("error")
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("cant start server:", err)
	}

}
