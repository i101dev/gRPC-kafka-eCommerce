// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: proto/product.proto

package proto

import (
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

// --------------------------------------------------------------------------
// Ping routes
type ProductPingOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingOrderReq) Reset() {
	*x = ProductPingOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingOrderReq) ProtoMessage() {}

func (x *ProductPingOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingOrderReq.ProtoReflect.Descriptor instead.
func (*ProductPingOrderReq) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductPingOrderReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProductPingOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingOrderRes) Reset() {
	*x = ProductPingOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingOrderRes) ProtoMessage() {}

func (x *ProductPingOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingOrderRes.ProtoReflect.Descriptor instead.
func (*ProductPingOrderRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductPingOrderRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProductPingUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingUserReq) Reset() {
	*x = ProductPingUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingUserReq) ProtoMessage() {}

func (x *ProductPingUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingUserReq.ProtoReflect.Descriptor instead.
func (*ProductPingUserReq) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductPingUserReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProductPingUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingUserRes) Reset() {
	*x = ProductPingUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingUserRes) ProtoMessage() {}

func (x *ProductPingUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingUserRes.ProtoReflect.Descriptor instead.
func (*ProductPingUserRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductPingUserRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProductPingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingReq) Reset() {
	*x = ProductPingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingReq) ProtoMessage() {}

func (x *ProductPingReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingReq.ProtoReflect.Descriptor instead.
func (*ProductPingReq) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductPingReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ProductPingRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductPingRes) Reset() {
	*x = ProductPingRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPingRes) ProtoMessage() {}

func (x *ProductPingRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPingRes.ProtoReflect.Descriptor instead.
func (*ProductPingRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductPingRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// --------------------------------------------------------------------------
// Connections
type ProductConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderSrv string `protobuf:"bytes,1,opt,name=orderSrv,proto3" json:"orderSrv,omitempty"`
	UserSrv  string `protobuf:"bytes,2,opt,name=userSrv,proto3" json:"userSrv,omitempty"`
}

func (x *ProductConnReq) Reset() {
	*x = ProductConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductConnReq) ProtoMessage() {}

func (x *ProductConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductConnReq.ProtoReflect.Descriptor instead.
func (*ProductConnReq) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductConnReq) GetOrderSrv() string {
	if x != nil {
		return x.OrderSrv
	}
	return ""
}

func (x *ProductConnReq) GetUserSrv() string {
	if x != nil {
		return x.UserSrv
	}
	return ""
}

type ProductConnRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductConnRes) Reset() {
	*x = ProductConnRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductConnRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductConnRes) ProtoMessage() {}

func (x *ProductConnRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductConnRes.ProtoReflect.Descriptor instead.
func (*ProductConnRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{7}
}

func (x *ProductConnRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// --------------------------------------------------------------------------
// Test
type ProductTestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Val int64  `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *ProductTestReq) Reset() {
	*x = ProductTestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductTestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductTestReq) ProtoMessage() {}

func (x *ProductTestReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductTestReq.ProtoReflect.Descriptor instead.
func (*ProductTestReq) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{8}
}

func (x *ProductTestReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ProductTestReq) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type ProductTestRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProductTestRes) Reset() {
	*x = ProductTestRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductTestRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductTestRes) ProtoMessage() {}

func (x *ProductTestRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductTestRes.ProtoReflect.Descriptor instead.
func (*ProductTestRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{9}
}

func (x *ProductTestRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_product_proto protoreflect.FileDescriptor

var file_proto_product_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x22, 0x27, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x27, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x26, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x26, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x22, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x46, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x72, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x72, 0x76, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x72, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x72, 0x76, 0x22, 0x22, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x34, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x22, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0x8e, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69,
	0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData = file_proto_product_proto_rawDesc
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_proto_rawDescData)
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_product_proto_goTypes = []interface{}{
	(*ProductPingOrderReq)(nil), // 0: ecommerce.ProductPingOrderReq
	(*ProductPingOrderRes)(nil), // 1: ecommerce.ProductPingOrderRes
	(*ProductPingUserReq)(nil),  // 2: ecommerce.ProductPingUserReq
	(*ProductPingUserRes)(nil),  // 3: ecommerce.ProductPingUserRes
	(*ProductPingReq)(nil),      // 4: ecommerce.ProductPingReq
	(*ProductPingRes)(nil),      // 5: ecommerce.ProductPingRes
	(*ProductConnReq)(nil),      // 6: ecommerce.ProductConnReq
	(*ProductConnRes)(nil),      // 7: ecommerce.ProductConnRes
	(*ProductTestReq)(nil),      // 8: ecommerce.ProductTestReq
	(*ProductTestRes)(nil),      // 9: ecommerce.ProductTestRes
}
var file_proto_product_proto_depIdxs = []int32{
	8, // 0: ecommerce.ProductService.ProductTest:input_type -> ecommerce.ProductTestReq
	6, // 1: ecommerce.ProductService.ProductConn:input_type -> ecommerce.ProductConnReq
	0, // 2: ecommerce.ProductService.ProductPingOrder:input_type -> ecommerce.ProductPingOrderReq
	2, // 3: ecommerce.ProductService.ProductPingUser:input_type -> ecommerce.ProductPingUserReq
	4, // 4: ecommerce.ProductService.ProductPing:input_type -> ecommerce.ProductPingReq
	9, // 5: ecommerce.ProductService.ProductTest:output_type -> ecommerce.ProductTestRes
	7, // 6: ecommerce.ProductService.ProductConn:output_type -> ecommerce.ProductConnRes
	1, // 7: ecommerce.ProductService.ProductPingOrder:output_type -> ecommerce.ProductPingOrderRes
	3, // 8: ecommerce.ProductService.ProductPingUser:output_type -> ecommerce.ProductPingUserRes
	5, // 9: ecommerce.ProductService.ProductPing:output_type -> ecommerce.ProductPingRes
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingOrderReq); i {
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
		file_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingOrderRes); i {
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
		file_proto_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingUserReq); i {
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
		file_proto_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingUserRes); i {
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
		file_proto_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingReq); i {
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
		file_proto_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPingRes); i {
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
		file_proto_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductConnReq); i {
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
		file_proto_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductConnRes); i {
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
		file_proto_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductTestReq); i {
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
		file_proto_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductTestRes); i {
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
			RawDescriptor: file_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_rawDesc = nil
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
}
