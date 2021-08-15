package main

import (
	"github.com/mahesh-dilhan/protogrpc/app1/core"
	grpc2 "github.com/mahesh-dilhan/protogrpc/app1/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	// configure our core service
	userService := core.NewService()
	// configure our gRPC service controller
	userServiceController := NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	grpc2.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)
	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
