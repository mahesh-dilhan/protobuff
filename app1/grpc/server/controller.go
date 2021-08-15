package main

import (
	"context"
	"github.com/mahesh-dilhan/protogrpc/app1"
	"github.com/mahesh-dilhan/protogrpc/app1/grpc"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService app1.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService app1.Service) grpc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *grpc.GetUsersRequest) (resp *grpc.GetUsersResponse, err error) {
	resultMap, err := ctlr.userService.GetUsers(req.GetIds())
	if err != nil {
		return
	}
	resp = &grpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *app1.User) *grpc.User {
	return &grpc.User{Id: u.ID, Name: u.Name}
}
