// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ServiceDiscoveryInitClient is the client API for ServiceDiscoveryInit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceDiscoveryInitClient interface {
	RegisterService(ctx context.Context, in *RegisterData, opts ...grpc.CallOption) (*ReturnPayload, error)
	DeleteService(ctx context.Context, in *ServiceGuid, opts ...grpc.CallOption) (*ReturnPayload, error)
	UpdateServiceHealth(ctx context.Context, in *RegisterData, opts ...grpc.CallOption) (*ReturnPayload, error)
}

type serviceDiscoveryInitClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceDiscoveryInitClient(cc grpc.ClientConnInterface) ServiceDiscoveryInitClient {
	return &serviceDiscoveryInitClient{cc}
}

func (c *serviceDiscoveryInitClient) RegisterService(ctx context.Context, in *RegisterData, opts ...grpc.CallOption) (*ReturnPayload, error) {
	out := new(ReturnPayload)
	err := c.cc.Invoke(ctx, "/protogen.ServiceDiscoveryInit/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceDiscoveryInitClient) DeleteService(ctx context.Context, in *ServiceGuid, opts ...grpc.CallOption) (*ReturnPayload, error) {
	out := new(ReturnPayload)
	err := c.cc.Invoke(ctx, "/protogen.ServiceDiscoveryInit/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceDiscoveryInitClient) UpdateServiceHealth(ctx context.Context, in *RegisterData, opts ...grpc.CallOption) (*ReturnPayload, error) {
	out := new(ReturnPayload)
	err := c.cc.Invoke(ctx, "/protogen.ServiceDiscoveryInit/UpdateServiceHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceDiscoveryInitServer is the server API for ServiceDiscoveryInit service.
// All implementations must embed UnimplementedServiceDiscoveryInitServer
// for forward compatibility
type ServiceDiscoveryInitServer interface {
	RegisterService(context.Context, *RegisterData) (*ReturnPayload, error)
	DeleteService(context.Context, *ServiceGuid) (*ReturnPayload, error)
	UpdateServiceHealth(context.Context, *RegisterData) (*ReturnPayload, error)
	mustEmbedUnimplementedServiceDiscoveryInitServer()
}

// UnimplementedServiceDiscoveryInitServer must be embedded to have forward compatible implementations.
type UnimplementedServiceDiscoveryInitServer struct {
}

func (UnimplementedServiceDiscoveryInitServer) RegisterService(context.Context, *RegisterData) (*ReturnPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedServiceDiscoveryInitServer) DeleteService(context.Context, *ServiceGuid) (*ReturnPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServiceDiscoveryInitServer) UpdateServiceHealth(context.Context, *RegisterData) (*ReturnPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceHealth not implemented")
}
func (UnimplementedServiceDiscoveryInitServer) mustEmbedUnimplementedServiceDiscoveryInitServer() {}

// UnsafeServiceDiscoveryInitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceDiscoveryInitServer will
// result in compilation errors.
type UnsafeServiceDiscoveryInitServer interface {
	mustEmbedUnimplementedServiceDiscoveryInitServer()
}

func RegisterServiceDiscoveryInitServer(s grpc.ServiceRegistrar, srv ServiceDiscoveryInitServer) {
	s.RegisterService(&ServiceDiscoveryInit_ServiceDesc, srv)
}

func _ServiceDiscoveryInit_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDiscoveryInitServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.ServiceDiscoveryInit/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDiscoveryInitServer).RegisterService(ctx, req.(*RegisterData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceDiscoveryInit_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceGuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDiscoveryInitServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.ServiceDiscoveryInit/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDiscoveryInitServer).DeleteService(ctx, req.(*ServiceGuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceDiscoveryInit_UpdateServiceHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDiscoveryInitServer).UpdateServiceHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.ServiceDiscoveryInit/UpdateServiceHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDiscoveryInitServer).UpdateServiceHealth(ctx, req.(*RegisterData))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceDiscoveryInit_ServiceDesc is the grpc.ServiceDesc for ServiceDiscoveryInit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceDiscoveryInit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protogen.ServiceDiscoveryInit",
	HandlerType: (*ServiceDiscoveryInitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _ServiceDiscoveryInit_RegisterService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServiceDiscoveryInit_DeleteService_Handler,
		},
		{
			MethodName: "UpdateServiceHealth",
			Handler:    _ServiceDiscoveryInit_UpdateServiceHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/servicediscovery.proto",
}

// ServiceDiscoveryInfoClient is the client API for ServiceDiscoveryInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceDiscoveryInfoClient interface {
	GetAllServices(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Services, error)
	GetByNameService(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*Services, error)
}

type serviceDiscoveryInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceDiscoveryInfoClient(cc grpc.ClientConnInterface) ServiceDiscoveryInfoClient {
	return &serviceDiscoveryInfoClient{cc}
}

func (c *serviceDiscoveryInfoClient) GetAllServices(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Services, error) {
	out := new(Services)
	err := c.cc.Invoke(ctx, "/protogen.ServiceDiscoveryInfo/GetAllServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceDiscoveryInfoClient) GetByNameService(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (*Services, error) {
	out := new(Services)
	err := c.cc.Invoke(ctx, "/protogen.ServiceDiscoveryInfo/GetByNameService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceDiscoveryInfoServer is the server API for ServiceDiscoveryInfo service.
// All implementations must embed UnimplementedServiceDiscoveryInfoServer
// for forward compatibility
type ServiceDiscoveryInfoServer interface {
	GetAllServices(context.Context, *EmptyRequest) (*Services, error)
	GetByNameService(context.Context, *ServiceName) (*Services, error)
	mustEmbedUnimplementedServiceDiscoveryInfoServer()
}

// UnimplementedServiceDiscoveryInfoServer must be embedded to have forward compatible implementations.
type UnimplementedServiceDiscoveryInfoServer struct {
}

func (UnimplementedServiceDiscoveryInfoServer) GetAllServices(context.Context, *EmptyRequest) (*Services, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllServices not implemented")
}
func (UnimplementedServiceDiscoveryInfoServer) GetByNameService(context.Context, *ServiceName) (*Services, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByNameService not implemented")
}
func (UnimplementedServiceDiscoveryInfoServer) mustEmbedUnimplementedServiceDiscoveryInfoServer() {}

// UnsafeServiceDiscoveryInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceDiscoveryInfoServer will
// result in compilation errors.
type UnsafeServiceDiscoveryInfoServer interface {
	mustEmbedUnimplementedServiceDiscoveryInfoServer()
}

func RegisterServiceDiscoveryInfoServer(s grpc.ServiceRegistrar, srv ServiceDiscoveryInfoServer) {
	s.RegisterService(&ServiceDiscoveryInfo_ServiceDesc, srv)
}

func _ServiceDiscoveryInfo_GetAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDiscoveryInfoServer).GetAllServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.ServiceDiscoveryInfo/GetAllServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDiscoveryInfoServer).GetAllServices(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceDiscoveryInfo_GetByNameService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDiscoveryInfoServer).GetByNameService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protogen.ServiceDiscoveryInfo/GetByNameService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDiscoveryInfoServer).GetByNameService(ctx, req.(*ServiceName))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceDiscoveryInfo_ServiceDesc is the grpc.ServiceDesc for ServiceDiscoveryInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceDiscoveryInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protogen.ServiceDiscoveryInfo",
	HandlerType: (*ServiceDiscoveryInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllServices",
			Handler:    _ServiceDiscoveryInfo_GetAllServices_Handler,
		},
		{
			MethodName: "GetByNameService",
			Handler:    _ServiceDiscoveryInfo_GetByNameService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/servicediscovery.proto",
}
