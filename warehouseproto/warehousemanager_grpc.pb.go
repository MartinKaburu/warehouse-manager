// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: warehousemanager.proto

package warehouseproto

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

// OrderConsumerClient is the client API for OrderConsumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderConsumerClient interface {
	// Receives a stream of orders
	ReceiveOrders(ctx context.Context, opts ...grpc.CallOption) (OrderConsumer_ReceiveOrdersClient, error)
}

type orderConsumerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderConsumerClient(cc grpc.ClientConnInterface) OrderConsumerClient {
	return &orderConsumerClient{cc}
}

func (c *orderConsumerClient) ReceiveOrders(ctx context.Context, opts ...grpc.CallOption) (OrderConsumer_ReceiveOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderConsumer_ServiceDesc.Streams[0], "/warehouseproto.OrderConsumer/ReceiveOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderConsumerReceiveOrdersClient{stream}
	return x, nil
}

type OrderConsumer_ReceiveOrdersClient interface {
	Send(*OrderRequest) error
	CloseAndRecv() (*OrderReport, error)
	grpc.ClientStream
}

type orderConsumerReceiveOrdersClient struct {
	grpc.ClientStream
}

func (x *orderConsumerReceiveOrdersClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderConsumerReceiveOrdersClient) CloseAndRecv() (*OrderReport, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OrderReport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderConsumerServer is the server API for OrderConsumer service.
// All implementations must embed UnimplementedOrderConsumerServer
// for forward compatibility
type OrderConsumerServer interface {
	// Receives a stream of orders
	ReceiveOrders(OrderConsumer_ReceiveOrdersServer) error
	mustEmbedUnimplementedOrderConsumerServer()
}

// UnimplementedOrderConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderConsumerServer struct {
}

func (UnimplementedOrderConsumerServer) ReceiveOrders(OrderConsumer_ReceiveOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveOrders not implemented")
}
func (UnimplementedOrderConsumerServer) mustEmbedUnimplementedOrderConsumerServer() {}

// UnsafeOrderConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderConsumerServer will
// result in compilation errors.
type UnsafeOrderConsumerServer interface {
	mustEmbedUnimplementedOrderConsumerServer()
}

func RegisterOrderConsumerServer(s grpc.ServiceRegistrar, srv OrderConsumerServer) {
	s.RegisterService(&OrderConsumer_ServiceDesc, srv)
}

func _OrderConsumer_ReceiveOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderConsumerServer).ReceiveOrders(&orderConsumerReceiveOrdersServer{stream})
}

type OrderConsumer_ReceiveOrdersServer interface {
	SendAndClose(*OrderReport) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderConsumerReceiveOrdersServer struct {
	grpc.ServerStream
}

func (x *orderConsumerReceiveOrdersServer) SendAndClose(m *OrderReport) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderConsumerReceiveOrdersServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderConsumer_ServiceDesc is the grpc.ServiceDesc for OrderConsumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderConsumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouseproto.OrderConsumer",
	HandlerType: (*OrderConsumerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveOrders",
			Handler:       _OrderConsumer_ReceiveOrders_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "warehousemanager.proto",
}
