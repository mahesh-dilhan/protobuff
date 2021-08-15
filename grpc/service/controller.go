package main

import (
	"context"
	"github.com/mahesh-dilhan/protogrpc"
	mysvcgrpc "github.com/mahesh-dilhan/protogrpc/grpc"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService protogrpc.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService protogrpc.Service) mysvcgrpc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *mysvcgrpc.GetUsersRequest) (resp *mysvcgrpc.GetUsersResponse, err error) {
	resultMap, err := ctlr.userService.GetUsers(req.GetIds())
	if err != nil {
		return
	}
	resp = &mysvcgrpc.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *protogrpc.User) *mysvcgrpc.User {
	return &mysvcgrpc.User{Id: u.ID, Name: u.Name}
}
