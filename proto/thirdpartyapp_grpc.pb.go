// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: thirdpartyapp.proto

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

// ThirdPartyAppServiceClient is the client API for ThirdPartyAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdPartyAppServiceClient interface {
	CreateThirdPartyApp(ctx context.Context, in *CreateThirdPartyAppRequest, opts ...grpc.CallOption) (*CreateThirdPartyAppResponse, error)
	ListThirdPartyApps(ctx context.Context, in *ListThirdPartyAppsRequest, opts ...grpc.CallOption) (*ListThirdPartyAppsResponse, error)
	ListEnterpriseThirdPartyApps(ctx context.Context, in *ListThirdPartyAppsRequest, opts ...grpc.CallOption) (*ListThirdPartyAppsResponse, error)
	ListThirdPartyTemplates(ctx context.Context, in *ListThirdPartyTemplatesRequest, opts ...grpc.CallOption) (*ListThirdPartyTemplatesResponse, error)
	GetThirdPartyApp(ctx context.Context, in *GetThirdPartyAppRequest, opts ...grpc.CallOption) (*GetThirdPartyAppResponse, error)
	DeleteThirdPartyApp(ctx context.Context, in *DeleteThirdPartyAppRequest, opts ...grpc.CallOption) (*DeleteThirdPartyAppResponse, error)
	UploadThirdPartyApp(ctx context.Context, in *UploadThirdPartyAppRequest, opts ...grpc.CallOption) (*UploadThirdPartyAppResponse, error)
	PreviewThirdPartyApp(ctx context.Context, in *PreviewThirdPartyAppRequest, opts ...grpc.CallOption) (*PreviewThirdPartyAppResponse, error)
	SubmitThirdPartyApp(ctx context.Context, in *SubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*SubmitThirdPartyAppResponse, error)
	BatchSubmitThirdPartyApp(ctx context.Context, in *BatchSubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*BatchSubmitThirdPartyAppResponse, error)
	ReleaseThirdPartyApp(ctx context.Context, in *ReleaseThirdPartyAppRequest, opts ...grpc.CallOption) (*ReleaseThirdPartyAppResponse, error)
	BatchReleaseThirdPartyApp(ctx context.Context, in *BatchSubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*BatchSubmitThirdPartyAppResponse, error)
	WithdrawThirdPartyApp(ctx context.Context, in *WithdrawThirdPartyAppRequest, opts ...grpc.CallOption) (*WithdrawThirdPartyAppResponse, error)
}

type thirdPartyAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyAppServiceClient(cc grpc.ClientConnInterface) ThirdPartyAppServiceClient {
	return &thirdPartyAppServiceClient{cc}
}

func (c *thirdPartyAppServiceClient) CreateThirdPartyApp(ctx context.Context, in *CreateThirdPartyAppRequest, opts ...grpc.CallOption) (*CreateThirdPartyAppResponse, error) {
	out := new(CreateThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/CreateThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) ListThirdPartyApps(ctx context.Context, in *ListThirdPartyAppsRequest, opts ...grpc.CallOption) (*ListThirdPartyAppsResponse, error) {
	out := new(ListThirdPartyAppsResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/ListThirdPartyApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) ListEnterpriseThirdPartyApps(ctx context.Context, in *ListThirdPartyAppsRequest, opts ...grpc.CallOption) (*ListThirdPartyAppsResponse, error) {
	out := new(ListThirdPartyAppsResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/ListEnterpriseThirdPartyApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) ListThirdPartyTemplates(ctx context.Context, in *ListThirdPartyTemplatesRequest, opts ...grpc.CallOption) (*ListThirdPartyTemplatesResponse, error) {
	out := new(ListThirdPartyTemplatesResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/ListThirdPartyTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) GetThirdPartyApp(ctx context.Context, in *GetThirdPartyAppRequest, opts ...grpc.CallOption) (*GetThirdPartyAppResponse, error) {
	out := new(GetThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/GetThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) DeleteThirdPartyApp(ctx context.Context, in *DeleteThirdPartyAppRequest, opts ...grpc.CallOption) (*DeleteThirdPartyAppResponse, error) {
	out := new(DeleteThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/DeleteThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) UploadThirdPartyApp(ctx context.Context, in *UploadThirdPartyAppRequest, opts ...grpc.CallOption) (*UploadThirdPartyAppResponse, error) {
	out := new(UploadThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/UploadThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) PreviewThirdPartyApp(ctx context.Context, in *PreviewThirdPartyAppRequest, opts ...grpc.CallOption) (*PreviewThirdPartyAppResponse, error) {
	out := new(PreviewThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/PreviewThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) SubmitThirdPartyApp(ctx context.Context, in *SubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*SubmitThirdPartyAppResponse, error) {
	out := new(SubmitThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/SubmitThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) BatchSubmitThirdPartyApp(ctx context.Context, in *BatchSubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*BatchSubmitThirdPartyAppResponse, error) {
	out := new(BatchSubmitThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/BatchSubmitThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) ReleaseThirdPartyApp(ctx context.Context, in *ReleaseThirdPartyAppRequest, opts ...grpc.CallOption) (*ReleaseThirdPartyAppResponse, error) {
	out := new(ReleaseThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/ReleaseThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) BatchReleaseThirdPartyApp(ctx context.Context, in *BatchSubmitThirdPartyAppRequest, opts ...grpc.CallOption) (*BatchSubmitThirdPartyAppResponse, error) {
	out := new(BatchSubmitThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/BatchReleaseThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyAppServiceClient) WithdrawThirdPartyApp(ctx context.Context, in *WithdrawThirdPartyAppRequest, opts ...grpc.CallOption) (*WithdrawThirdPartyAppResponse, error) {
	out := new(WithdrawThirdPartyAppResponse)
	err := c.cc.Invoke(ctx, "/thirdpartyapp.ThirdPartyAppService/WithdrawThirdPartyApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyAppServiceServer is the server API for ThirdPartyAppService service.
// All implementations must embed UnimplementedThirdPartyAppServiceServer
// for forward compatibility
type ThirdPartyAppServiceServer interface {
	CreateThirdPartyApp(context.Context, *CreateThirdPartyAppRequest) (*CreateThirdPartyAppResponse, error)
	ListThirdPartyApps(context.Context, *ListThirdPartyAppsRequest) (*ListThirdPartyAppsResponse, error)
	ListEnterpriseThirdPartyApps(context.Context, *ListThirdPartyAppsRequest) (*ListThirdPartyAppsResponse, error)
	ListThirdPartyTemplates(context.Context, *ListThirdPartyTemplatesRequest) (*ListThirdPartyTemplatesResponse, error)
	GetThirdPartyApp(context.Context, *GetThirdPartyAppRequest) (*GetThirdPartyAppResponse, error)
	DeleteThirdPartyApp(context.Context, *DeleteThirdPartyAppRequest) (*DeleteThirdPartyAppResponse, error)
	UploadThirdPartyApp(context.Context, *UploadThirdPartyAppRequest) (*UploadThirdPartyAppResponse, error)
	PreviewThirdPartyApp(context.Context, *PreviewThirdPartyAppRequest) (*PreviewThirdPartyAppResponse, error)
	SubmitThirdPartyApp(context.Context, *SubmitThirdPartyAppRequest) (*SubmitThirdPartyAppResponse, error)
	BatchSubmitThirdPartyApp(context.Context, *BatchSubmitThirdPartyAppRequest) (*BatchSubmitThirdPartyAppResponse, error)
	ReleaseThirdPartyApp(context.Context, *ReleaseThirdPartyAppRequest) (*ReleaseThirdPartyAppResponse, error)
	BatchReleaseThirdPartyApp(context.Context, *BatchSubmitThirdPartyAppRequest) (*BatchSubmitThirdPartyAppResponse, error)
	WithdrawThirdPartyApp(context.Context, *WithdrawThirdPartyAppRequest) (*WithdrawThirdPartyAppResponse, error)
	mustEmbedUnimplementedThirdPartyAppServiceServer()
}

// UnimplementedThirdPartyAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThirdPartyAppServiceServer struct {
}

func (UnimplementedThirdPartyAppServiceServer) CreateThirdPartyApp(context.Context, *CreateThirdPartyAppRequest) (*CreateThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) ListThirdPartyApps(context.Context, *ListThirdPartyAppsRequest) (*ListThirdPartyAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThirdPartyApps not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) ListEnterpriseThirdPartyApps(context.Context, *ListThirdPartyAppsRequest) (*ListThirdPartyAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnterpriseThirdPartyApps not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) ListThirdPartyTemplates(context.Context, *ListThirdPartyTemplatesRequest) (*ListThirdPartyTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListThirdPartyTemplates not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) GetThirdPartyApp(context.Context, *GetThirdPartyAppRequest) (*GetThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) DeleteThirdPartyApp(context.Context, *DeleteThirdPartyAppRequest) (*DeleteThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) UploadThirdPartyApp(context.Context, *UploadThirdPartyAppRequest) (*UploadThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) PreviewThirdPartyApp(context.Context, *PreviewThirdPartyAppRequest) (*PreviewThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) SubmitThirdPartyApp(context.Context, *SubmitThirdPartyAppRequest) (*SubmitThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) BatchSubmitThirdPartyApp(context.Context, *BatchSubmitThirdPartyAppRequest) (*BatchSubmitThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSubmitThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) ReleaseThirdPartyApp(context.Context, *ReleaseThirdPartyAppRequest) (*ReleaseThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) BatchReleaseThirdPartyApp(context.Context, *BatchSubmitThirdPartyAppRequest) (*BatchSubmitThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchReleaseThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) WithdrawThirdPartyApp(context.Context, *WithdrawThirdPartyAppRequest) (*WithdrawThirdPartyAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawThirdPartyApp not implemented")
}
func (UnimplementedThirdPartyAppServiceServer) mustEmbedUnimplementedThirdPartyAppServiceServer() {}

// UnsafeThirdPartyAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdPartyAppServiceServer will
// result in compilation errors.
type UnsafeThirdPartyAppServiceServer interface {
	mustEmbedUnimplementedThirdPartyAppServiceServer()
}

func RegisterThirdPartyAppServiceServer(s grpc.ServiceRegistrar, srv ThirdPartyAppServiceServer) {
	s.RegisterService(&ThirdPartyAppService_ServiceDesc, srv)
}

func _ThirdPartyAppService_CreateThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).CreateThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/CreateThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).CreateThirdPartyApp(ctx, req.(*CreateThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_ListThirdPartyApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListThirdPartyAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).ListThirdPartyApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/ListThirdPartyApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).ListThirdPartyApps(ctx, req.(*ListThirdPartyAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_ListEnterpriseThirdPartyApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListThirdPartyAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).ListEnterpriseThirdPartyApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/ListEnterpriseThirdPartyApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).ListEnterpriseThirdPartyApps(ctx, req.(*ListThirdPartyAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_ListThirdPartyTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListThirdPartyTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).ListThirdPartyTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/ListThirdPartyTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).ListThirdPartyTemplates(ctx, req.(*ListThirdPartyTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_GetThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).GetThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/GetThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).GetThirdPartyApp(ctx, req.(*GetThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_DeleteThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).DeleteThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/DeleteThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).DeleteThirdPartyApp(ctx, req.(*DeleteThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_UploadThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).UploadThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/UploadThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).UploadThirdPartyApp(ctx, req.(*UploadThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_PreviewThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).PreviewThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/PreviewThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).PreviewThirdPartyApp(ctx, req.(*PreviewThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_SubmitThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).SubmitThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/SubmitThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).SubmitThirdPartyApp(ctx, req.(*SubmitThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_BatchSubmitThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSubmitThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).BatchSubmitThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/BatchSubmitThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).BatchSubmitThirdPartyApp(ctx, req.(*BatchSubmitThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_ReleaseThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).ReleaseThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/ReleaseThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).ReleaseThirdPartyApp(ctx, req.(*ReleaseThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_BatchReleaseThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSubmitThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).BatchReleaseThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/BatchReleaseThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).BatchReleaseThirdPartyApp(ctx, req.(*BatchSubmitThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyAppService_WithdrawThirdPartyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawThirdPartyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppServiceServer).WithdrawThirdPartyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdpartyapp.ThirdPartyAppService/WithdrawThirdPartyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppServiceServer).WithdrawThirdPartyApp(ctx, req.(*WithdrawThirdPartyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdPartyAppService_ServiceDesc is the grpc.ServiceDesc for ThirdPartyAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdPartyAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thirdpartyapp.ThirdPartyAppService",
	HandlerType: (*ThirdPartyAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateThirdPartyApp",
			Handler:    _ThirdPartyAppService_CreateThirdPartyApp_Handler,
		},
		{
			MethodName: "ListThirdPartyApps",
			Handler:    _ThirdPartyAppService_ListThirdPartyApps_Handler,
		},
		{
			MethodName: "ListEnterpriseThirdPartyApps",
			Handler:    _ThirdPartyAppService_ListEnterpriseThirdPartyApps_Handler,
		},
		{
			MethodName: "ListThirdPartyTemplates",
			Handler:    _ThirdPartyAppService_ListThirdPartyTemplates_Handler,
		},
		{
			MethodName: "GetThirdPartyApp",
			Handler:    _ThirdPartyAppService_GetThirdPartyApp_Handler,
		},
		{
			MethodName: "DeleteThirdPartyApp",
			Handler:    _ThirdPartyAppService_DeleteThirdPartyApp_Handler,
		},
		{
			MethodName: "UploadThirdPartyApp",
			Handler:    _ThirdPartyAppService_UploadThirdPartyApp_Handler,
		},
		{
			MethodName: "PreviewThirdPartyApp",
			Handler:    _ThirdPartyAppService_PreviewThirdPartyApp_Handler,
		},
		{
			MethodName: "SubmitThirdPartyApp",
			Handler:    _ThirdPartyAppService_SubmitThirdPartyApp_Handler,
		},
		{
			MethodName: "BatchSubmitThirdPartyApp",
			Handler:    _ThirdPartyAppService_BatchSubmitThirdPartyApp_Handler,
		},
		{
			MethodName: "ReleaseThirdPartyApp",
			Handler:    _ThirdPartyAppService_ReleaseThirdPartyApp_Handler,
		},
		{
			MethodName: "BatchReleaseThirdPartyApp",
			Handler:    _ThirdPartyAppService_BatchReleaseThirdPartyApp_Handler,
		},
		{
			MethodName: "WithdrawThirdPartyApp",
			Handler:    _ThirdPartyAppService_WithdrawThirdPartyApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thirdpartyapp.proto",
}
