// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: pkg/pb/travel.proto

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

// TravelServiceClient is the client API for TravelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelServiceClient interface {
	GetsTravel(ctx context.Context, in *GetsTravelRequest, opts ...grpc.CallOption) (*GetsTravelResponse, error)
	GetTravel(ctx context.Context, in *GetTravelRequest, opts ...grpc.CallOption) (*GetTravelResponse, error)
	CreateTravel(ctx context.Context, in *CreateTravelRequest, opts ...grpc.CallOption) (*CreateTravelResponse, error)
	UpdateTravel(ctx context.Context, in *UpdateTravelRequest, opts ...grpc.CallOption) (*UpdateTravelResponse, error)
	DeleteTravel(ctx context.Context, in *DeleteTravelRequest, opts ...grpc.CallOption) (*DeleteTravelResponse, error)
}

type travelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelServiceClient(cc grpc.ClientConnInterface) TravelServiceClient {
	return &travelServiceClient{cc}
}

func (c *travelServiceClient) GetsTravel(ctx context.Context, in *GetsTravelRequest, opts ...grpc.CallOption) (*GetsTravelResponse, error) {
	out := new(GetsTravelResponse)
	err := c.cc.Invoke(ctx, "/travel.TravelService/GetsTravel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelServiceClient) GetTravel(ctx context.Context, in *GetTravelRequest, opts ...grpc.CallOption) (*GetTravelResponse, error) {
	out := new(GetTravelResponse)
	err := c.cc.Invoke(ctx, "/travel.TravelService/GetTravel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelServiceClient) CreateTravel(ctx context.Context, in *CreateTravelRequest, opts ...grpc.CallOption) (*CreateTravelResponse, error) {
	out := new(CreateTravelResponse)
	err := c.cc.Invoke(ctx, "/travel.TravelService/CreateTravel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelServiceClient) UpdateTravel(ctx context.Context, in *UpdateTravelRequest, opts ...grpc.CallOption) (*UpdateTravelResponse, error) {
	out := new(UpdateTravelResponse)
	err := c.cc.Invoke(ctx, "/travel.TravelService/UpdateTravel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelServiceClient) DeleteTravel(ctx context.Context, in *DeleteTravelRequest, opts ...grpc.CallOption) (*DeleteTravelResponse, error) {
	out := new(DeleteTravelResponse)
	err := c.cc.Invoke(ctx, "/travel.TravelService/DeleteTravel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelServiceServer is the server API for TravelService service.
// All implementations should embed UnimplementedTravelServiceServer
// for forward compatibility
type TravelServiceServer interface {
	GetsTravel(context.Context, *GetsTravelRequest) (*GetsTravelResponse, error)
	GetTravel(context.Context, *GetTravelRequest) (*GetTravelResponse, error)
	CreateTravel(context.Context, *CreateTravelRequest) (*CreateTravelResponse, error)
	UpdateTravel(context.Context, *UpdateTravelRequest) (*UpdateTravelResponse, error)
	DeleteTravel(context.Context, *DeleteTravelRequest) (*DeleteTravelResponse, error)
}

// UnimplementedTravelServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTravelServiceServer struct {
}

func (UnimplementedTravelServiceServer) GetsTravel(context.Context, *GetsTravelRequest) (*GetsTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsTravel not implemented")
}
func (UnimplementedTravelServiceServer) GetTravel(context.Context, *GetTravelRequest) (*GetTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTravel not implemented")
}
func (UnimplementedTravelServiceServer) CreateTravel(context.Context, *CreateTravelRequest) (*CreateTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTravel not implemented")
}
func (UnimplementedTravelServiceServer) UpdateTravel(context.Context, *UpdateTravelRequest) (*UpdateTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTravel not implemented")
}
func (UnimplementedTravelServiceServer) DeleteTravel(context.Context, *DeleteTravelRequest) (*DeleteTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTravel not implemented")
}

// UnsafeTravelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelServiceServer will
// result in compilation errors.
type UnsafeTravelServiceServer interface {
	mustEmbedUnimplementedTravelServiceServer()
}

func RegisterTravelServiceServer(s grpc.ServiceRegistrar, srv TravelServiceServer) {
	s.RegisterService(&TravelService_ServiceDesc, srv)
}

func _TravelService_GetsTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServiceServer).GetsTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/travel.TravelService/GetsTravel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServiceServer).GetsTravel(ctx, req.(*GetsTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelService_GetTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServiceServer).GetTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/travel.TravelService/GetTravel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServiceServer).GetTravel(ctx, req.(*GetTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelService_CreateTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServiceServer).CreateTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/travel.TravelService/CreateTravel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServiceServer).CreateTravel(ctx, req.(*CreateTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelService_UpdateTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServiceServer).UpdateTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/travel.TravelService/UpdateTravel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServiceServer).UpdateTravel(ctx, req.(*UpdateTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelService_DeleteTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServiceServer).DeleteTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/travel.TravelService/DeleteTravel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServiceServer).DeleteTravel(ctx, req.(*DeleteTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TravelService_ServiceDesc is the grpc.ServiceDesc for TravelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TravelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "travel.TravelService",
	HandlerType: (*TravelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetsTravel",
			Handler:    _TravelService_GetsTravel_Handler,
		},
		{
			MethodName: "GetTravel",
			Handler:    _TravelService_GetTravel_Handler,
		},
		{
			MethodName: "CreateTravel",
			Handler:    _TravelService_CreateTravel_Handler,
		},
		{
			MethodName: "UpdateTravel",
			Handler:    _TravelService_UpdateTravel_Handler,
		},
		{
			MethodName: "DeleteTravel",
			Handler:    _TravelService_DeleteTravel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/travel.proto",
}
