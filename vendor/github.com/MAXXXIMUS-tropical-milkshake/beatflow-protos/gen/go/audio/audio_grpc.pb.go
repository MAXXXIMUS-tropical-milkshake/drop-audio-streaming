// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: audio/audio.proto

package audiov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AudioService_Upload_FullMethodName     = "/audio.AudioService/Upload"
	AudioService_GetBeat_FullMethodName    = "/audio.AudioService/GetBeat"
	AudioService_GetFilters_FullMethodName = "/audio.AudioService/GetFilters"
)

// AudioServiceClient is the client API for AudioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioServiceClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	GetBeat(ctx context.Context, in *GetBeatRequest, opts ...grpc.CallOption) (*GetBeatResponse, error)
	GetFilters(ctx context.Context, in *GetFiltersRequest, opts ...grpc.CallOption) (*GetFiltersResponse, error)
}

type audioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioServiceClient(cc grpc.ClientConnInterface) AudioServiceClient {
	return &audioServiceClient{cc}
}

func (c *audioServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, AudioService_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetBeat(ctx context.Context, in *GetBeatRequest, opts ...grpc.CallOption) (*GetBeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBeatResponse)
	err := c.cc.Invoke(ctx, AudioService_GetBeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audioServiceClient) GetFilters(ctx context.Context, in *GetFiltersRequest, opts ...grpc.CallOption) (*GetFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFiltersResponse)
	err := c.cc.Invoke(ctx, AudioService_GetFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioServiceServer is the server API for AudioService service.
// All implementations must embed UnimplementedAudioServiceServer
// for forward compatibility.
type AudioServiceServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	GetBeat(context.Context, *GetBeatRequest) (*GetBeatResponse, error)
	GetFilters(context.Context, *GetFiltersRequest) (*GetFiltersResponse, error)
	mustEmbedUnimplementedAudioServiceServer()
}

// UnimplementedAudioServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAudioServiceServer struct{}

func (UnimplementedAudioServiceServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedAudioServiceServer) GetBeat(context.Context, *GetBeatRequest) (*GetBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeat not implemented")
}
func (UnimplementedAudioServiceServer) GetFilters(context.Context, *GetFiltersRequest) (*GetFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilters not implemented")
}
func (UnimplementedAudioServiceServer) mustEmbedUnimplementedAudioServiceServer() {}
func (UnimplementedAudioServiceServer) testEmbeddedByValue()                      {}

// UnsafeAudioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioServiceServer will
// result in compilation errors.
type UnsafeAudioServiceServer interface {
	mustEmbedUnimplementedAudioServiceServer()
}

func RegisterAudioServiceServer(s grpc.ServiceRegistrar, srv AudioServiceServer) {
	// If the following call pancis, it indicates UnimplementedAudioServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AudioService_ServiceDesc, srv)
}

func _AudioService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetBeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetBeat(ctx, req.(*GetBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudioService_GetFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioServiceServer).GetFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudioService_GetFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioServiceServer).GetFilters(ctx, req.(*GetFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioService_ServiceDesc is the grpc.ServiceDesc for AudioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audio.AudioService",
	HandlerType: (*AudioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _AudioService_Upload_Handler,
		},
		{
			MethodName: "GetBeat",
			Handler:    _AudioService_GetBeat_Handler,
		},
		{
			MethodName: "GetFilters",
			Handler:    _AudioService_GetFilters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audio/audio.proto",
}
