// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: libcore.proto

package gen

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
	LibcoreService_Exit_FullMethodName                = "/libcore.LibcoreService/Exit"
	LibcoreService_Update_FullMethodName              = "/libcore.LibcoreService/Update"
	LibcoreService_Start_FullMethodName               = "/libcore.LibcoreService/Start"
	LibcoreService_Stop_FullMethodName                = "/libcore.LibcoreService/Stop"
	LibcoreService_Test_FullMethodName                = "/libcore.LibcoreService/Test"
	LibcoreService_QueryStats_FullMethodName          = "/libcore.LibcoreService/QueryStats"
	LibcoreService_ListConnections_FullMethodName     = "/libcore.LibcoreService/ListConnections"
	LibcoreService_GetGeoIPList_FullMethodName        = "/libcore.LibcoreService/GetGeoIPList"
	LibcoreService_GetGeoSiteList_FullMethodName      = "/libcore.LibcoreService/GetGeoSiteList"
	LibcoreService_CompileGeoIPToSrs_FullMethodName   = "/libcore.LibcoreService/CompileGeoIPToSrs"
	LibcoreService_CompileGeoSiteToSrs_FullMethodName = "/libcore.LibcoreService/CompileGeoSiteToSrs"
	LibcoreService_SetSystemProxy_FullMethodName      = "/libcore.LibcoreService/SetSystemProxy"
)

// LibcoreServiceClient is the client API for LibcoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibcoreServiceClient interface {
	Exit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	Start(ctx context.Context, in *LoadConfigReq, opts ...grpc.CallOption) (*ErrorResp, error)
	Stop(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ErrorResp, error)
	Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error)
	QueryStats(ctx context.Context, in *QueryStatsReq, opts ...grpc.CallOption) (*QueryStatsResp, error)
	ListConnections(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListConnectionsResp, error)
	GetGeoIPList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetGeoIPListResponse, error)
	GetGeoSiteList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetGeoSiteListResponse, error)
	CompileGeoIPToSrs(ctx context.Context, in *CompileGeoIPToSrsRequest, opts ...grpc.CallOption) (*EmptyResp, error)
	CompileGeoSiteToSrs(ctx context.Context, in *CompileGeoSiteToSrsRequest, opts ...grpc.CallOption) (*EmptyResp, error)
	SetSystemProxy(ctx context.Context, in *SetSystemProxyRequest, opts ...grpc.CallOption) (*EmptyResp, error)
}

type libcoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibcoreServiceClient(cc grpc.ClientConnInterface) LibcoreServiceClient {
	return &libcoreServiceClient{cc}
}

func (c *libcoreServiceClient) Exit(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, LibcoreService_Exit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, LibcoreService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Start(ctx context.Context, in *LoadConfigReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	out := new(ErrorResp)
	err := c.cc.Invoke(ctx, LibcoreService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Stop(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	out := new(ErrorResp)
	err := c.cc.Invoke(ctx, LibcoreService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error) {
	out := new(TestResp)
	err := c.cc.Invoke(ctx, LibcoreService_Test_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) QueryStats(ctx context.Context, in *QueryStatsReq, opts ...grpc.CallOption) (*QueryStatsResp, error) {
	out := new(QueryStatsResp)
	err := c.cc.Invoke(ctx, LibcoreService_QueryStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) ListConnections(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListConnectionsResp, error) {
	out := new(ListConnectionsResp)
	err := c.cc.Invoke(ctx, LibcoreService_ListConnections_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) GetGeoIPList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetGeoIPListResponse, error) {
	out := new(GetGeoIPListResponse)
	err := c.cc.Invoke(ctx, LibcoreService_GetGeoIPList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) GetGeoSiteList(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetGeoSiteListResponse, error) {
	out := new(GetGeoSiteListResponse)
	err := c.cc.Invoke(ctx, LibcoreService_GetGeoSiteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) CompileGeoIPToSrs(ctx context.Context, in *CompileGeoIPToSrsRequest, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, LibcoreService_CompileGeoIPToSrs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) CompileGeoSiteToSrs(ctx context.Context, in *CompileGeoSiteToSrsRequest, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, LibcoreService_CompileGeoSiteToSrs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcoreServiceClient) SetSystemProxy(ctx context.Context, in *SetSystemProxyRequest, opts ...grpc.CallOption) (*EmptyResp, error) {
	out := new(EmptyResp)
	err := c.cc.Invoke(ctx, LibcoreService_SetSystemProxy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibcoreServiceServer is the server API for LibcoreService service.
// All implementations must embed UnimplementedLibcoreServiceServer
// for forward compatibility
type LibcoreServiceServer interface {
	Exit(context.Context, *EmptyReq) (*EmptyResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	Start(context.Context, *LoadConfigReq) (*ErrorResp, error)
	Stop(context.Context, *EmptyReq) (*ErrorResp, error)
	Test(context.Context, *TestReq) (*TestResp, error)
	QueryStats(context.Context, *QueryStatsReq) (*QueryStatsResp, error)
	ListConnections(context.Context, *EmptyReq) (*ListConnectionsResp, error)
	GetGeoIPList(context.Context, *EmptyReq) (*GetGeoIPListResponse, error)
	GetGeoSiteList(context.Context, *EmptyReq) (*GetGeoSiteListResponse, error)
	CompileGeoIPToSrs(context.Context, *CompileGeoIPToSrsRequest) (*EmptyResp, error)
	CompileGeoSiteToSrs(context.Context, *CompileGeoSiteToSrsRequest) (*EmptyResp, error)
	SetSystemProxy(context.Context, *SetSystemProxyRequest) (*EmptyResp, error)
	mustEmbedUnimplementedLibcoreServiceServer()
}

// UnimplementedLibcoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibcoreServiceServer struct {
}

func (UnimplementedLibcoreServiceServer) Exit(context.Context, *EmptyReq) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedLibcoreServiceServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLibcoreServiceServer) Start(context.Context, *LoadConfigReq) (*ErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedLibcoreServiceServer) Stop(context.Context, *EmptyReq) (*ErrorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedLibcoreServiceServer) Test(context.Context, *TestReq) (*TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedLibcoreServiceServer) QueryStats(context.Context, *QueryStatsReq) (*QueryStatsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}
func (UnimplementedLibcoreServiceServer) ListConnections(context.Context, *EmptyReq) (*ListConnectionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnections not implemented")
}
func (UnimplementedLibcoreServiceServer) GetGeoIPList(context.Context, *EmptyReq) (*GetGeoIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoIPList not implemented")
}
func (UnimplementedLibcoreServiceServer) GetGeoSiteList(context.Context, *EmptyReq) (*GetGeoSiteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeoSiteList not implemented")
}
func (UnimplementedLibcoreServiceServer) CompileGeoIPToSrs(context.Context, *CompileGeoIPToSrsRequest) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompileGeoIPToSrs not implemented")
}
func (UnimplementedLibcoreServiceServer) CompileGeoSiteToSrs(context.Context, *CompileGeoSiteToSrsRequest) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompileGeoSiteToSrs not implemented")
}
func (UnimplementedLibcoreServiceServer) SetSystemProxy(context.Context, *SetSystemProxyRequest) (*EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemProxy not implemented")
}
func (UnimplementedLibcoreServiceServer) mustEmbedUnimplementedLibcoreServiceServer() {}

// UnsafeLibcoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibcoreServiceServer will
// result in compilation errors.
type UnsafeLibcoreServiceServer interface {
	mustEmbedUnimplementedLibcoreServiceServer()
}

func RegisterLibcoreServiceServer(s grpc.ServiceRegistrar, srv LibcoreServiceServer) {
	s.RegisterService(&LibcoreService_ServiceDesc, srv)
}

func _LibcoreService_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_Exit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Exit(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Start(ctx, req.(*LoadConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Stop(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).Test(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_QueryStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).QueryStats(ctx, req.(*QueryStatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_ListConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).ListConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_ListConnections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).ListConnections(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_GetGeoIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).GetGeoIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_GetGeoIPList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).GetGeoIPList(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_GetGeoSiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).GetGeoSiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_GetGeoSiteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).GetGeoSiteList(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_CompileGeoIPToSrs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileGeoIPToSrsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).CompileGeoIPToSrs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_CompileGeoIPToSrs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).CompileGeoIPToSrs(ctx, req.(*CompileGeoIPToSrsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_CompileGeoSiteToSrs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileGeoSiteToSrsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).CompileGeoSiteToSrs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_CompileGeoSiteToSrs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).CompileGeoSiteToSrs(ctx, req.(*CompileGeoSiteToSrsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcoreService_SetSystemProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSystemProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcoreServiceServer).SetSystemProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibcoreService_SetSystemProxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcoreServiceServer).SetSystemProxy(ctx, req.(*SetSystemProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibcoreService_ServiceDesc is the grpc.ServiceDesc for LibcoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibcoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libcore.LibcoreService",
	HandlerType: (*LibcoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exit",
			Handler:    _LibcoreService_Exit_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LibcoreService_Update_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _LibcoreService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LibcoreService_Stop_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _LibcoreService_Test_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _LibcoreService_QueryStats_Handler,
		},
		{
			MethodName: "ListConnections",
			Handler:    _LibcoreService_ListConnections_Handler,
		},
		{
			MethodName: "GetGeoIPList",
			Handler:    _LibcoreService_GetGeoIPList_Handler,
		},
		{
			MethodName: "GetGeoSiteList",
			Handler:    _LibcoreService_GetGeoSiteList_Handler,
		},
		{
			MethodName: "CompileGeoIPToSrs",
			Handler:    _LibcoreService_CompileGeoIPToSrs_Handler,
		},
		{
			MethodName: "CompileGeoSiteToSrs",
			Handler:    _LibcoreService_CompileGeoSiteToSrs_Handler,
		},
		{
			MethodName: "SetSystemProxy",
			Handler:    _LibcoreService_SetSystemProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libcore.proto",
}
