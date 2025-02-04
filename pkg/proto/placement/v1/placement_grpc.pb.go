// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package placement

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

// PlacementClient is the client API for Placement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlacementClient interface {
	ReportDaprStatus(ctx context.Context, opts ...grpc.CallOption) (Placement_ReportDaprStatusClient, error)
}

type placementClient struct {
	cc grpc.ClientConnInterface
}

func NewPlacementClient(cc grpc.ClientConnInterface) PlacementClient {
	return &placementClient{cc}
}

func (c *placementClient) ReportDaprStatus(ctx context.Context, opts ...grpc.CallOption) (Placement_ReportDaprStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Placement_ServiceDesc.Streams[0], "/dapr.proto.placement.v1.Placement/ReportDaprStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &placementReportDaprStatusClient{stream}
	return x, nil
}

type Placement_ReportDaprStatusClient interface {
	Send(*Host) error
	Recv() (*PlacementOrder, error)
	grpc.ClientStream
}

type placementReportDaprStatusClient struct {
	grpc.ClientStream
}

func (x *placementReportDaprStatusClient) Send(m *Host) error {
	return x.ClientStream.SendMsg(m)
}

func (x *placementReportDaprStatusClient) Recv() (*PlacementOrder, error) {
	m := new(PlacementOrder)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlacementServer is the server API for Placement service.
// All implementations should embed UnimplementedPlacementServer
// for forward compatibility
type PlacementServer interface {
	ReportDaprStatus(Placement_ReportDaprStatusServer) error
}

// UnimplementedPlacementServer should be embedded to have forward compatible implementations.
type UnimplementedPlacementServer struct {
}

func (UnimplementedPlacementServer) ReportDaprStatus(Placement_ReportDaprStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportDaprStatus not implemented")
}

// UnsafePlacementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlacementServer will
// result in compilation errors.
type UnsafePlacementServer interface {
	mustEmbedUnimplementedPlacementServer()
}

func RegisterPlacementServer(s grpc.ServiceRegistrar, srv PlacementServer) {
	s.RegisterService(&Placement_ServiceDesc, srv)
}

func _Placement_ReportDaprStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlacementServer).ReportDaprStatus(&placementReportDaprStatusServer{stream})
}

type Placement_ReportDaprStatusServer interface {
	Send(*PlacementOrder) error
	Recv() (*Host, error)
	grpc.ServerStream
}

type placementReportDaprStatusServer struct {
	grpc.ServerStream
}

func (x *placementReportDaprStatusServer) Send(m *PlacementOrder) error {
	return x.ServerStream.SendMsg(m)
}

func (x *placementReportDaprStatusServer) Recv() (*Host, error) {
	m := new(Host)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Placement_ServiceDesc is the grpc.ServiceDesc for Placement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Placement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.placement.v1.Placement",
	HandlerType: (*PlacementServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportDaprStatus",
			Handler:       _Placement_ReportDaprStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dapr/proto/placement/v1/placement.proto",
}
