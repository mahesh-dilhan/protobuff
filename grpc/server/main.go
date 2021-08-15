package main

import (
	protogrpccore "github.com/mahesh-dilhan/protogrpc/core"
	protogrpcgrpc "github.com/mahesh-dilhan/protogrpc/grpc"
	protogrpcgrpcservice "github.com/mahesh-dilhan/protogrpc/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	// configure our core service
	userService := protogrpccore.NewService()
	// configure our gRPC service controller
	userServiceController := protogrpcgrpcservice.NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	protogrpcgrpc.RegisterUserServiceServer(server, userServiceController)
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
