// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

var userCreateUserStreamDesc = &grpc.StreamDesc{
	StreamName: "CreateUser",
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.pbnj.api.proto.v1.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userDeleteUserStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteUser",
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.pbnj.api.proto.v1.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var userUpdateUserStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdateUser",
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.pbnj.api.proto.v1.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserService is the service API for User service.
// Fields should be assigned to their respective handler implementations only before
// RegisterUserService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type UserService struct {
	CreateUser func(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser func(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	UpdateUser func(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func (s *UserService) createUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/github.com.tinkerbell.pbnj.api.proto.v1.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserService) deleteUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/github.com.tinkerbell.pbnj.api.proto.v1.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *UserService) updateUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/github.com.tinkerbell.pbnj.api.proto.v1.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterUserService registers a service implementation with a gRPC server.
func RegisterUserService(s grpc.ServiceRegistrar, srv *UserService) {
	srvCopy := *srv
	if srvCopy.CreateUser == nil {
		srvCopy.CreateUser = func(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
		}
	}
	if srvCopy.DeleteUser == nil {
		srvCopy.DeleteUser = func(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
		}
	}
	if srvCopy.UpdateUser == nil {
		srvCopy.UpdateUser = func(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "github.com.tinkerbell.pbnj.api.proto.v1.User",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CreateUser",
				Handler:    srvCopy.createUser,
			},
			{
				MethodName: "DeleteUser",
				Handler:    srvCopy.deleteUser,
			},
			{
				MethodName: "UpdateUser",
				Handler:    srvCopy.updateUser,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "api/proto/v1/user.proto",
	}

	s.RegisterService(&sd, nil)
}
