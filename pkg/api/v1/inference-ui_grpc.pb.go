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

// InferenceUIClient is the client API for InferenceUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InferenceUIClient interface {
	// ListEngineSpecs returns a list of Inference Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (InferenceUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type inferenceUIClient struct {
	cc grpc.ClientConnInterface
}

func NewInferenceUIClient(cc grpc.ClientConnInterface) InferenceUIClient {
	return &inferenceUIClient{cc}
}

func (c *inferenceUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (InferenceUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InferenceUI_ServiceDesc.Streams[0], "/v1.InferenceUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &inferenceUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InferenceUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type inferenceUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *inferenceUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *inferenceUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.InferenceUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InferenceUIServer is the server API for InferenceUI service.
// All implementations must embed UnimplementedInferenceUIServer
// for forward compatibility
type InferenceUIServer interface {
	// ListEngineSpecs returns a list of Inference Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, InferenceUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedInferenceUIServer()
}

// UnimplementedInferenceUIServer must be embedded to have forward compatible implementations.
type UnimplementedInferenceUIServer struct {
}

func (UnimplementedInferenceUIServer) ListEngineSpecs(*ListEngineSpecsRequest, InferenceUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedInferenceUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedInferenceUIServer) mustEmbedUnimplementedInferenceUIServer() {}

// UnsafeInferenceUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InferenceUIServer will
// result in compilation errors.
type UnsafeInferenceUIServer interface {
	mustEmbedUnimplementedInferenceUIServer()
}

func RegisterInferenceUIServer(s grpc.ServiceRegistrar, srv InferenceUIServer) {
	s.RegisterService(&InferenceUI_ServiceDesc, srv)
}

func _InferenceUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InferenceUIServer).ListEngineSpecs(m, &inferenceUIListEngineSpecsServer{stream})
}

type InferenceUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type inferenceUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *inferenceUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InferenceUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InferenceUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.InferenceUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InferenceUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InferenceUI_ServiceDesc is the grpc.ServiceDesc for InferenceUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InferenceUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.InferenceUI",
	HandlerType: (*InferenceUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _InferenceUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _InferenceUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inference-ui.proto",
}
