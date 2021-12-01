// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Proto

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

// AuctionHouseClient is the client API for AuctionHouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionHouseClient interface {
	Bid(ctx context.Context, in *Offer, opts ...grpc.CallOption) (*Ack, error)
	Result(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Details, error)
}

type auctionHouseClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionHouseClient(cc grpc.ClientConnInterface) AuctionHouseClient {
	return &auctionHouseClient{cc}
}

func (c *auctionHouseClient) Bid(ctx context.Context, in *Offer, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Proto.AuctionHouse/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Result(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Details, error) {
	out := new(Details)
	err := c.cc.Invoke(ctx, "/Proto.AuctionHouse/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionHouseServer is the server API for AuctionHouse service.
// All implementations must embed UnimplementedAuctionHouseServer
// for forward compatibility
type AuctionHouseServer interface {
	Bid(context.Context, *Offer) (*Ack, error)
	Result(context.Context, *Info) (*Details, error)
	mustEmbedUnimplementedAuctionHouseServer()
}

// UnimplementedAuctionHouseServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionHouseServer struct {
}

func (UnimplementedAuctionHouseServer) Bid(context.Context, *Offer) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionHouseServer) Result(context.Context, *Info) (*Details, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionHouseServer) mustEmbedUnimplementedAuctionHouseServer() {}

// UnsafeAuctionHouseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionHouseServer will
// result in compilation errors.
type UnsafeAuctionHouseServer interface {
	mustEmbedUnimplementedAuctionHouseServer()
}

func RegisterAuctionHouseServer(s grpc.ServiceRegistrar, srv AuctionHouseServer) {
	s.RegisterService(&AuctionHouse_ServiceDesc, srv)
}

func _AuctionHouse_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Offer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.AuctionHouse/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Bid(ctx, req.(*Offer))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.AuctionHouse/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Result(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionHouse_ServiceDesc is the grpc.ServiceDesc for AuctionHouse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionHouse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.AuctionHouse",
	HandlerType: (*AuctionHouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _AuctionHouse_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionHouse_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ah.proto",
}
