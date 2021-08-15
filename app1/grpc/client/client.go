package main

import (
	"context"
	"fmt"
	"github.com/mahesh-dilhan/protogrpc/app1"
	grpc2 "github.com/mahesh-dilhan/protogrpc/app1/grpc"
	"google.golang.org/grpc"
	"time"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient grpc2.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (app1.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: grpc2.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers(ids []int64) (result map[int64]app1.User, err error) {
	result = map[int64]app1.User{}
	req := &grpc2.GetUsersRequest{
		Ids: ids,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		u := unmarshalUser(grpcUser)
		result[u.ID] = u
	}
	return
}
func (s *grpcService) GetUser(id int64) (result app1.User, err error) {
	req := &grpc2.GetUsersRequest{
		Ids: []int64{id},
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		// sanity check: only the requested user should be present in results
		if grpcUser.GetId() == id {
			return unmarshalUser(grpcUser), nil
		}
	}
	return result, app1.ErrNotFound
}
func unmarshalUser(grpcUser *grpc2.User) (result app1.User) {
	result.ID = grpcUser.Id
	result.Name = grpcUser.Name
	return
}

func main() {
	client, err := NewGRPCService("127.0.0.1:9000")
	if err != nil {
		fmt.Println("Error connecting")
	}
	fmt.Println(client)
	fmt.Println(client.GetUser(1))
}
