// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/friends.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FriendsServiceClient is the client API for FriendsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendsServiceClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*Response, error)
	AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*Response, error)
	RejectFriend(ctx context.Context, in *RejectFriendRequest, opts ...grpc.CallOption) (*Response, error)
	GetAllFriendRequests(ctx context.Context, in *GetAllFriendRequestsRequest, opts ...grpc.CallOption) (*GetAllFriendRequestsResponse, error)
	GetAllFriends(ctx context.Context, in *GetAllFriendsRequest, opts ...grpc.CallOption) (*GetAllFriendsResponse, error)
}

type friendsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendsServiceClient(cc grpc.ClientConnInterface) FriendsServiceClient {
	return &friendsServiceClient{cc}
}

func (c *friendsServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/friendspb.FriendsService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/friendspb.FriendsService/AcceptFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) RejectFriend(ctx context.Context, in *RejectFriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/friendspb.FriendsService/RejectFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) GetAllFriendRequests(ctx context.Context, in *GetAllFriendRequestsRequest, opts ...grpc.CallOption) (*GetAllFriendRequestsResponse, error) {
	out := new(GetAllFriendRequestsResponse)
	err := c.cc.Invoke(ctx, "/friendspb.FriendsService/GetAllFriendRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) GetAllFriends(ctx context.Context, in *GetAllFriendsRequest, opts ...grpc.CallOption) (*GetAllFriendsResponse, error) {
	out := new(GetAllFriendsResponse)
	err := c.cc.Invoke(ctx, "/friendspb.FriendsService/GetAllFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendsServiceServer is the server API for FriendsService service.
// All implementations must embed UnimplementedFriendsServiceServer
// for forward compatibility
type FriendsServiceServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*Response, error)
	AcceptFriend(context.Context, *AcceptFriendRequest) (*Response, error)
	RejectFriend(context.Context, *RejectFriendRequest) (*Response, error)
	GetAllFriendRequests(context.Context, *GetAllFriendRequestsRequest) (*GetAllFriendRequestsResponse, error)
	GetAllFriends(context.Context, *GetAllFriendsRequest) (*GetAllFriendsResponse, error)
	mustEmbedUnimplementedFriendsServiceServer()
}

// UnimplementedFriendsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendsServiceServer struct {
}

func (UnimplementedFriendsServiceServer) AddFriend(context.Context, *AddFriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendsServiceServer) AcceptFriend(context.Context, *AcceptFriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriend not implemented")
}
func (UnimplementedFriendsServiceServer) RejectFriend(context.Context, *RejectFriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectFriend not implemented")
}
func (UnimplementedFriendsServiceServer) GetAllFriendRequests(context.Context, *GetAllFriendRequestsRequest) (*GetAllFriendRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFriendRequests not implemented")
}
func (UnimplementedFriendsServiceServer) GetAllFriends(context.Context, *GetAllFriendsRequest) (*GetAllFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFriends not implemented")
}
func (UnimplementedFriendsServiceServer) mustEmbedUnimplementedFriendsServiceServer() {}

// UnsafeFriendsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendsServiceServer will
// result in compilation errors.
type UnsafeFriendsServiceServer interface {
	mustEmbedUnimplementedFriendsServiceServer()
}

func RegisterFriendsServiceServer(s grpc.ServiceRegistrar, srv FriendsServiceServer) {
	s.RegisterService(&FriendsService_ServiceDesc, srv)
}

func _FriendsService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friendspb.FriendsService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_AcceptFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).AcceptFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friendspb.FriendsService/AcceptFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).AcceptFriend(ctx, req.(*AcceptFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_RejectFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).RejectFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friendspb.FriendsService/RejectFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).RejectFriend(ctx, req.(*RejectFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_GetAllFriendRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFriendRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).GetAllFriendRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friendspb.FriendsService/GetAllFriendRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).GetAllFriendRequests(ctx, req.(*GetAllFriendRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_GetAllFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).GetAllFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friendspb.FriendsService/GetAllFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).GetAllFriends(ctx, req.(*GetAllFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendsService_ServiceDesc is the grpc.ServiceDesc for FriendsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friendspb.FriendsService",
	HandlerType: (*FriendsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _FriendsService_AddFriend_Handler,
		},
		{
			MethodName: "AcceptFriend",
			Handler:    _FriendsService_AcceptFriend_Handler,
		},
		{
			MethodName: "RejectFriend",
			Handler:    _FriendsService_RejectFriend_Handler,
		},
		{
			MethodName: "GetAllFriendRequests",
			Handler:    _FriendsService_GetAllFriendRequests_Handler,
		},
		{
			MethodName: "GetAllFriends",
			Handler:    _FriendsService_GetAllFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/friends.proto",
}
