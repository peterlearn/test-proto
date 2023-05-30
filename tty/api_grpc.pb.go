// 定义项目 API 的 proto 文件 可以同时描述 gRPC 和 HTTP API
// protobuf 文件参考:
//  - https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: tty/api.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

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

const (
	Tty_TtyVersion_FullMethodName = "/tty.Tty/TtyVersion"
)

// TtyClient is the client API for Tty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TtyClient interface {
	// 增加玩家购买的礼包次数
	TtyVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TtyVersionResp, error)
}

type ttyClient struct {
	cc grpc.ClientConnInterface
}

func NewTtyClient(cc grpc.ClientConnInterface) TtyClient {
	return &ttyClient{cc}
}

func (c *ttyClient) TtyVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TtyVersionResp, error) {
	out := new(TtyVersionResp)
	err := c.cc.Invoke(ctx, Tty_TtyVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TtyServer is the server API for Tty service.
// All implementations must embed UnimplementedTtyServer
// for forward compatibility
type TtyServer interface {
	// 增加玩家购买的礼包次数
	TtyVersion(context.Context, *Empty) (*TtyVersionResp, error)
	mustEmbedUnimplementedTtyServer()
}

// UnimplementedTtyServer must be embedded to have forward compatible implementations.
type UnimplementedTtyServer struct {
}

func (UnimplementedTtyServer) TtyVersion(context.Context, *Empty) (*TtyVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TtyVersion not implemented")
}
func (UnimplementedTtyServer) mustEmbedUnimplementedTtyServer() {}

// UnsafeTtyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TtyServer will
// result in compilation errors.
type UnsafeTtyServer interface {
	mustEmbedUnimplementedTtyServer()
}

func RegisterTtyServer(s grpc.ServiceRegistrar, srv TtyServer) {
	s.RegisterService(&Tty_ServiceDesc, srv)
}

func _Tty_TtyVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TtyServer).TtyVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tty_TtyVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TtyServer).TtyVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Tty_ServiceDesc is the grpc.ServiceDesc for Tty service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tty_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tty.Tty",
	HandlerType: (*TtyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TtyVersion",
			Handler:    _Tty_TtyVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tty/api.proto",
}