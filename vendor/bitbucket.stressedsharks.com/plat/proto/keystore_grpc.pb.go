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

// KeystoreClient is the client API for Keystore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeystoreClient interface {
	// Returns an extended private key.
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error)
	//Creates a new private key with the given alias
	CreateAlias(ctx context.Context, in *CreateAliasRequest, opts ...grpc.CallOption) (*CreateAliasResponse, error)
}

type keystoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKeystoreClient(cc grpc.ClientConnInterface) KeystoreClient {
	return &keystoreClient{cc}
}

func (c *keystoreClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error) {
	out := new(GetPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/proto.Keystore/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreClient) CreateAlias(ctx context.Context, in *CreateAliasRequest, opts ...grpc.CallOption) (*CreateAliasResponse, error) {
	out := new(CreateAliasResponse)
	err := c.cc.Invoke(ctx, "/proto.Keystore/CreateAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeystoreServer is the server API for Keystore service.
// All implementations must embed UnimplementedKeystoreServer
// for forward compatibility
type KeystoreServer interface {
	// Returns an extended private key.
	GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error)
	//Creates a new private key with the given alias
	CreateAlias(context.Context, *CreateAliasRequest) (*CreateAliasResponse, error)
	mustEmbedUnimplementedKeystoreServer()
}

// UnimplementedKeystoreServer must be embedded to have forward compatible implementations.
type UnimplementedKeystoreServer struct {
}

func (UnimplementedKeystoreServer) GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedKeystoreServer) CreateAlias(context.Context, *CreateAliasRequest) (*CreateAliasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlias not implemented")
}
func (UnimplementedKeystoreServer) mustEmbedUnimplementedKeystoreServer() {}

// UnsafeKeystoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeystoreServer will
// result in compilation errors.
type UnsafeKeystoreServer interface {
	mustEmbedUnimplementedKeystoreServer()
}

func RegisterKeystoreServer(s grpc.ServiceRegistrar, srv KeystoreServer) {
	s.RegisterService(&Keystore_ServiceDesc, srv)
}

func _Keystore_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Keystore/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keystore_CreateAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAliasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServer).CreateAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Keystore/CreateAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServer).CreateAlias(ctx, req.(*CreateAliasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Keystore_ServiceDesc is the grpc.ServiceDesc for Keystore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keystore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Keystore",
	HandlerType: (*KeystoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrivateKey",
			Handler:    _Keystore_GetPrivateKey_Handler,
		},
		{
			MethodName: "CreateAlias",
			Handler:    _Keystore_CreateAlias_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keystore.proto",
}
