// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: card.proto

package server

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

const (
	CardService_Play_FullMethodName  = "/server.CardService/Play"
	CardService_Draw_FullMethodName  = "/server.CardService/Draw"
	CardService_Stand_FullMethodName = "/server.CardService/Stand"
)

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceClient interface {
	Play(ctx context.Context, in *PlayCardRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Draw(ctx context.Context, in *DrawCardRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Stand(ctx context.Context, in *StandRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) Play(ctx context.Context, in *PlayCardRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, CardService_Play_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) Draw(ctx context.Context, in *DrawCardRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, CardService_Draw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) Stand(ctx context.Context, in *StandRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, CardService_Stand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
// All implementations must embed UnimplementedCardServiceServer
// for forward compatibility
type CardServiceServer interface {
	Play(context.Context, *PlayCardRequest) (*SuccessResponse, error)
	Draw(context.Context, *DrawCardRequest) (*SuccessResponse, error)
	Stand(context.Context, *StandRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedCardServiceServer()
}

// UnimplementedCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (UnimplementedCardServiceServer) Play(context.Context, *PlayCardRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedCardServiceServer) Draw(context.Context, *DrawCardRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Draw not implemented")
}
func (UnimplementedCardServiceServer) Stand(context.Context, *StandRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stand not implemented")
}
func (UnimplementedCardServiceServer) mustEmbedUnimplementedCardServiceServer() {}

// UnsafeCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceServer will
// result in compilation errors.
type UnsafeCardServiceServer interface {
	mustEmbedUnimplementedCardServiceServer()
}

func RegisterCardServiceServer(s grpc.ServiceRegistrar, srv CardServiceServer) {
	s.RegisterService(&CardService_ServiceDesc, srv)
}

func _CardService_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardService_Play_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).Play(ctx, req.(*PlayCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_Draw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).Draw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardService_Draw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).Draw(ctx, req.(*DrawCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_Stand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).Stand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardService_Stand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).Stand(ctx, req.(*StandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardService_ServiceDesc is the grpc.ServiceDesc for CardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _CardService_Play_Handler,
		},
		{
			MethodName: "Draw",
			Handler:    _CardService_Draw_Handler,
		},
		{
			MethodName: "Stand",
			Handler:    _CardService_Stand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card.proto",
}