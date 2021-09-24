// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: fund.proto

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

//*
//A Fund is a simple structure that holds information about a UTXO. This is used by the TransactionBuilder, SplittingService and FundingService.
type Fund struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid          string      `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`                                        // The transaction ID of the UTXO.
	Index         uint32      `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`                                     // The vout / n / index of the UTXO.
	LockingScript []byte      `protobuf:"bytes,3,opt,name=locking_script,json=lockingScript,proto3" json:"locking_script,omitempty"` // The Locking Script of the UTXO.
	Satoshis      uint64      `protobuf:"varint,4,opt,name=satoshis,proto3" json:"satoshis,omitempty"`                               // The value of the UTXO.
	Signer        *KeyContext `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *Fund) Reset() {
	*x = Fund{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fund) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fund) ProtoMessage() {}

func (x *Fund) ProtoReflect() protoreflect.Message {
	mi := &file_fund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fund.ProtoReflect.Descriptor instead.
func (*Fund) Descriptor() ([]byte, []int) {
	return file_fund_proto_rawDescGZIP(), []int{0}
}

func (x *Fund) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Fund) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Fund) GetLockingScript() []byte {
	if x != nil {
		return x.LockingScript
	}
	return nil
}

func (x *Fund) GetSatoshis() uint64 {
	if x != nil {
		return x.Satoshis
	}
	return 0
}

func (x *Fund) GetSigner() *KeyContext {
	if x != nil {
		return x.Signer
	}
	return nil
}

var File_fund_proto protoreflect.FileDescriptor

var file_fund_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x78, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x62, 0x69, 0x74, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x73, 0x68, 0x61,
	0x72, 0x6b, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fund_proto_rawDescOnce sync.Once
	file_fund_proto_rawDescData = file_fund_proto_rawDesc
)

func file_fund_proto_rawDescGZIP() []byte {
	file_fund_proto_rawDescOnce.Do(func() {
		file_fund_proto_rawDescData = protoimpl.X.CompressGZIP(file_fund_proto_rawDescData)
	})
	return file_fund_proto_rawDescData
}

var file_fund_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fund_proto_goTypes = []interface{}{
	(*Fund)(nil),       // 0: proto.Fund
	(*KeyContext)(nil), // 1: proto.KeyContext
}
var file_fund_proto_depIdxs = []int32{
	1, // 0: proto.Fund.signer:type_name -> proto.KeyContext
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fund_proto_init() }
func file_fund_proto_init() {
	if File_fund_proto != nil {
		return
	}
	file_key_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fund); i {
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
			RawDescriptor: file_fund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fund_proto_goTypes,
		DependencyIndexes: file_fund_proto_depIdxs,
		MessageInfos:      file_fund_proto_msgTypes,
	}.Build()
	File_fund_proto = out.File
	file_fund_proto_rawDesc = nil
	file_fund_proto_goTypes = nil
	file_fund_proto_depIdxs = nil
}
