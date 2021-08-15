package client

import (
	"context"
	"fmt"
	"github.com/mahesh-dilhan/protogrpc/app2/rpc"
	"google.golang.org/grpc"
	"os"
)

func clientMain() {
	if err := runClient(); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}
}
func runClient() error {
	// connnect
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}
	cache := rpc.NewCacheClient(conn)
	// store
	_, err = cache.Store(context.Background(), &rpc.StoreReq{Key: "gopher", Val: []byte("con")})
	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}
	// get
	resp, err := cache.Get(context.Background(), &rpc.GetReq{Key: "gopher"})
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}
	fmt.Printf("Got cached value %s\n", resp.Val)
	return nil
}
