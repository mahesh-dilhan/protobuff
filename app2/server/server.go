package server

import (
	"context"
	"fmt"
	"github.com/mahesh-dilhan/protogrpc/app2/rpc"
	"google.golang.org/grpc"
	"net"
	"os"
)

func serverMain() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
		os.Exit(1)
	}
}
func runServer() error {
	srv := grpc.NewServer()
	rpc.RegisterCacheServer(srv, &CacheService{})
	l, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		return err
	}
	// blocks until complete
	return srv.Serve(l)
}

type CacheService struct {
}

func (s *CacheService) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	return nil, fmt.Errorf("unimplemented")
}
func (s *CacheService) Store(ctx context.Context, req *rpc.StoreReq) (*rpc.StoreResp,
	error) {
	return nil, fmt.Errorf("unimplemented")
}
