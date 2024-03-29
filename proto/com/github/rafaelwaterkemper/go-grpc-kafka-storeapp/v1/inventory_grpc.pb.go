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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	// Check availability
	CheckAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityReply, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) CheckAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (*AvailabilityReply, error) {
	out := new(AvailabilityReply)
	err := c.cc.Invoke(ctx, "/inventory.Inventory/CheckAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	// Check availability
	CheckAvailability(context.Context, *AvailabilityRequest) (*AvailabilityReply, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) CheckAvailability(context.Context, *AvailabilityRequest) (*AvailabilityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAvailability not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Inventory/CheckAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CheckAvailability(ctx, req.(*AvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAvailability",
			Handler:    _Inventory_CheckAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/inventory.proto",
}
