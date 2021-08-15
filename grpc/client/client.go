package client

import (
	"context"
	"github.com/mahesh-dilhan/protogrpc"
	mysvcgrpc "github.com/mahesh-dilhan/protogrpc/grpc"
	"google.golang.org/grpc"
	"time"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient mysvcgrpc.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (protogrpc.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: mysvcgrpc.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers(ids []int64) (result map[int64]protogrpc.User, err error) {
	result = map[int64]protogrpc.User{}
	req := &mysvcgrpc.GetUsersRequest{
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
func (s *grpcService) GetUser(id int64) (result protogrpc.User, err error) {
	req := &mysvcgrpc.GetUsersRequest{
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
	return result, protogrpc.ErrNotFound
}
func unmarshalUser(grpcUser *mysvcgrpc.User) (result protogrpc.User) {
	result.ID = grpcUser.Id
	result.Name = grpcUser.Name
	return
}
