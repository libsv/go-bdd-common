// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: meta_map.proto

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

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyContext *KeyContext `protobuf:"bytes,1,opt,name=key_context,json=keyContext,proto3" json:"key_context,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetKeyContext() *KeyContext {
	if x != nil {
		return x.KeyContext
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentTxid string `protobuf:"bytes,1,opt,name=parent_txid,json=parentTxid,proto3" json:"parent_txid,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetParentTxid() string {
	if x != nil {
		return x.ParentTxid
	}
	return ""
}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyContext *KeyContext `protobuf:"bytes,1,opt,name=key_context,json=keyContext,proto3" json:"key_context,omitempty"`
	Txid       string      `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"` // The transaction ID of the transaction that was successfully written to the bitcoin node.
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{2}
}

func (x *StoreRequest) GetKeyContext() *KeyContext {
	if x != nil {
		return x.KeyContext
	}
	return nil
}

func (x *StoreRequest) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

// An empty response
type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{3}
}

type MarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid        string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`                                   // The transaction ID of the transaction that was successfully written to the bitcoin node.
	BlockHeight uint32 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"` // Optional - the block height that this transaction was seen in.
}

func (x *MarkRequest) Reset() {
	*x = MarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkRequest) ProtoMessage() {}

func (x *MarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkRequest.ProtoReflect.Descriptor instead.
func (*MarkRequest) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{4}
}

func (x *MarkRequest) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *MarkRequest) GetBlockHeight() uint32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

// An empty response
type MarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkResponse) Reset() {
	*x = MarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkResponse) ProtoMessage() {}

func (x *MarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkResponse.ProtoReflect.Descriptor instead.
func (*MarkResponse) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{5}
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyContext  *KeyContext `protobuf:"bytes,1,opt,name=key_context,json=keyContext,proto3" json:"key_context,omitempty"`
	IncludeData bool        `protobuf:"varint,2,opt,name=include_data,json=includeData,proto3" json:"include_data,omitempty"` // If true, the actual data for this transaction will be returned
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{6}
}

func (x *ReadRequest) GetKeyContext() *KeyContext {
	if x != nil {
		return x.KeyContext
	}
	return nil
}

func (x *ReadRequest) GetIncludeData() bool {
	if x != nil {
		return x.IncludeData
	}
	return false
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid        string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Marked      bool   `protobuf:"varint,2,opt,name=marked,proto3" json:"marked,omitempty"`                              // True indicates that this transaction has been seen on the blockchain
	BlockHeight uint32 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"` // The block height that this transaction is in, if it has been mined.
	Data        []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                                   // The data for this transaction.
	Timestamp   uint64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                        // The timestamp of this message
	Path        string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{7}
}

func (x *ReadResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *ReadResponse) GetMarked() bool {
	if x != nil {
		return x.Marked
	}
	return false
}

func (x *ReadResponse) GetBlockHeight() uint32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *ReadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReadResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ReadResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias         string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`                                       // The alias that represents the private key to use.
	LastTimestamp uint64 `protobuf:"varint,2,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"` // The timestamp of the last successfully received ReadResponse.
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_map_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meta_map_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_meta_map_proto_rawDescGZIP(), []int{8}
}

func (x *SubscribeRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SubscribeRequest) GetLastTimestamp() uint64 {
	if x != nil {
		return x.LastTimestamp
	}
	return 0
}

var File_meta_map_proto protoreflect.FileDescriptor

var file_meta_map_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x6b, 0x65,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x30,
	0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x78, 0x69, 0x64,
	0x22, 0x56, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x0b, 0x4d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x64, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x4f, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x9a, 0x02, 0x0a,
	0x07, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x62, 0x69, 0x74,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x73,
	0x68, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_map_proto_rawDescOnce sync.Once
	file_meta_map_proto_rawDescData = file_meta_map_proto_rawDesc
)

func file_meta_map_proto_rawDescGZIP() []byte {
	file_meta_map_proto_rawDescOnce.Do(func() {
		file_meta_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_map_proto_rawDescData)
	})
	return file_meta_map_proto_rawDescData
}

var file_meta_map_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_meta_map_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),     // 0: proto.CheckRequest
	(*CheckResponse)(nil),    // 1: proto.CheckResponse
	(*StoreRequest)(nil),     // 2: proto.StoreRequest
	(*StoreResponse)(nil),    // 3: proto.StoreResponse
	(*MarkRequest)(nil),      // 4: proto.MarkRequest
	(*MarkResponse)(nil),     // 5: proto.MarkResponse
	(*ReadRequest)(nil),      // 6: proto.ReadRequest
	(*ReadResponse)(nil),     // 7: proto.ReadResponse
	(*SubscribeRequest)(nil), // 8: proto.SubscribeRequest
	(*KeyContext)(nil),       // 9: proto.KeyContext
}
var file_meta_map_proto_depIdxs = []int32{
	9, // 0: proto.CheckRequest.key_context:type_name -> proto.KeyContext
	9, // 1: proto.StoreRequest.key_context:type_name -> proto.KeyContext
	9, // 2: proto.ReadRequest.key_context:type_name -> proto.KeyContext
	0, // 3: proto.MetaMap.Check:input_type -> proto.CheckRequest
	2, // 4: proto.MetaMap.Store:input_type -> proto.StoreRequest
	4, // 5: proto.MetaMap.Mark:input_type -> proto.MarkRequest
	6, // 6: proto.MetaMap.Read:input_type -> proto.ReadRequest
	8, // 7: proto.MetaMap.Subscribe:input_type -> proto.SubscribeRequest
	1, // 8: proto.MetaMap.Check:output_type -> proto.CheckResponse
	3, // 9: proto.MetaMap.Store:output_type -> proto.StoreResponse
	5, // 10: proto.MetaMap.Mark:output_type -> proto.MarkResponse
	7, // 11: proto.MetaMap.Read:output_type -> proto.ReadResponse
	7, // 12: proto.MetaMap.Subscribe:output_type -> proto.ReadResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_meta_map_proto_init() }
func file_meta_map_proto_init() {
	if File_meta_map_proto != nil {
		return
	}
	file_key_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meta_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_meta_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_meta_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_meta_map_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
		file_meta_map_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkRequest); i {
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
		file_meta_map_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkResponse); i {
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
		file_meta_map_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_meta_map_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_meta_map_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
			RawDescriptor: file_meta_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_map_proto_goTypes,
		DependencyIndexes: file_meta_map_proto_depIdxs,
		MessageInfos:      file_meta_map_proto_msgTypes,
	}.Build()
	File_meta_map_proto = out.File
	file_meta_map_proto_rawDesc = nil
	file_meta_map_proto_goTypes = nil
	file_meta_map_proto_depIdxs = nil
}
