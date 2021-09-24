// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: keystore.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetPrivateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"` // The alias that represents the private key to use (must be less than 256 characters).
}

func (x *GetPrivateKeyRequest) Reset() {
	*x = GetPrivateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateKeyRequest) ProtoMessage() {}

func (x *GetPrivateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keystore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateKeyRequest.ProtoReflect.Descriptor instead.
func (*GetPrivateKeyRequest) Descriptor() ([]byte, []int) {
	return file_keystore_proto_rawDescGZIP(), []int{0}
}

func (x *GetPrivateKeyRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type GetPrivateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtendedPrivateKey string `protobuf:"bytes,1,opt,name=extended_private_key,json=extendedPrivateKey,proto3" json:"extended_private_key,omitempty"` // The extended private key in xpriv format (will be exactly 111 characters).
}

func (x *GetPrivateKeyResponse) Reset() {
	*x = GetPrivateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateKeyResponse) ProtoMessage() {}

func (x *GetPrivateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keystore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateKeyResponse.ProtoReflect.Descriptor instead.
func (*GetPrivateKeyResponse) Descriptor() ([]byte, []int) {
	return file_keystore_proto_rawDescGZIP(), []int{1}
}

func (x *GetPrivateKeyResponse) GetExtendedPrivateKey() string {
	if x != nil {
		return x.ExtendedPrivateKey
	}
	return ""
}

type CreateAliasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"` // The alias associated with the private key
}

func (x *CreateAliasRequest) Reset() {
	*x = CreateAliasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAliasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAliasRequest) ProtoMessage() {}

func (x *CreateAliasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keystore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAliasRequest.ProtoReflect.Descriptor instead.
func (*CreateAliasRequest) Descriptor() ([]byte, []int) {
	return file_keystore_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAliasRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type CreateAliasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAliasResponse) Reset() {
	*x = CreateAliasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAliasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAliasResponse) ProtoMessage() {}

func (x *CreateAliasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keystore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAliasResponse.ProtoReflect.Descriptor instead.
func (*CreateAliasResponse) Descriptor() ([]byte, []int) {
	return file_keystore_proto_rawDescGZIP(), []int{3}
}

var File_keystore_proto protoreflect.FileDescriptor

var file_keystore_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x22, 0x2a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x15, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x73, 0x68, 0x61, 0x72,
	0x6b, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keystore_proto_rawDescOnce sync.Once
	file_keystore_proto_rawDescData = file_keystore_proto_rawDesc
)

func file_keystore_proto_rawDescGZIP() []byte {
	file_keystore_proto_rawDescOnce.Do(func() {
		file_keystore_proto_rawDescData = protoimpl.X.CompressGZIP(file_keystore_proto_rawDescData)
	})
	return file_keystore_proto_rawDescData
}

var file_keystore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_keystore_proto_goTypes = []interface{}{
	(*GetPrivateKeyRequest)(nil),  // 0: proto.GetPrivateKeyRequest
	(*GetPrivateKeyResponse)(nil), // 1: proto.GetPrivateKeyResponse
	(*CreateAliasRequest)(nil),    // 2: proto.CreateAliasRequest
	(*CreateAliasResponse)(nil),   // 3: proto.CreateAliasResponse
}
var file_keystore_proto_depIdxs = []int32{
	0, // 0: proto.Keystore.GetPrivateKey:input_type -> proto.GetPrivateKeyRequest
	2, // 1: proto.Keystore.CreateAlias:input_type -> proto.CreateAliasRequest
	1, // 2: proto.Keystore.GetPrivateKey:output_type -> proto.GetPrivateKeyResponse
	3, // 3: proto.Keystore.CreateAlias:output_type -> proto.CreateAliasResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keystore_proto_init() }
func file_keystore_proto_init() {
	if File_keystore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keystore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateKeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAliasRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keystore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAliasResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keystore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keystore_proto_goTypes,
		DependencyIndexes: file_keystore_proto_depIdxs,
		MessageInfos:      file_keystore_proto_msgTypes,
	}.Build()
	File_keystore_proto = out.File
	file_keystore_proto_rawDesc = nil
	file_keystore_proto_goTypes = nil
	file_keystore_proto_depIdxs = nil
}
