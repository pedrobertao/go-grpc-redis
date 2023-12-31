// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.3
// source: proto/go-grpc-redis.proto

package goredispb

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

// Result status code
type ResultStatus int32

const (
	ResultStatus_UNKNOWN  ResultStatus = 0
	ResultStatus_SUCCESS  ResultStatus = 1
	ResultStatus_FAIL     ResultStatus = 2
	ResultStatus_NOTFOUND ResultStatus = 3
	ResultStatus_UPDATED  ResultStatus = 4
)

// Enum value maps for ResultStatus.
var (
	ResultStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAIL",
		3: "NOTFOUND",
		4: "UPDATED",
	}
	ResultStatus_value = map[string]int32{
		"UNKNOWN":  0,
		"SUCCESS":  1,
		"FAIL":     2,
		"NOTFOUND": 3,
		"UPDATED":  4,
	}
)

func (x ResultStatus) Enum() *ResultStatus {
	p := new(ResultStatus)
	*p = x
	return p
}

func (x ResultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_go_grpc_redis_proto_enumTypes[0].Descriptor()
}

func (ResultStatus) Type() protoreflect.EnumType {
	return &file_proto_go_grpc_redis_proto_enumTypes[0]
}

func (x ResultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultStatus.Descriptor instead.
func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{0}
}

// HTTP result message
type HttpMessage int32

const (
	HttpMessage_UNSPECIFIED                     HttpMessage = 0
	HttpMessage_PROCESSING                      HttpMessage = 1
	HttpMessage_OK                              HttpMessage = 2
	HttpMessage_CREATED                         HttpMessage = 3
	HttpMessage_NO_CONTENT                      HttpMessage = 4
	HttpMessage_PARTIAL_CONTENT                 HttpMessage = 5
	HttpMessage_FOUND                           HttpMessage = 6
	HttpMessage_BAD_REQUEST                     HttpMessage = 7
	HttpMessage_UNAUTHORIZED                    HttpMessage = 8
	HttpMessage_FORBIDDEN                       HttpMessage = 9
	HttpMessage_NOT_FOUND                       HttpMessage = 10
	HttpMessage_METHOD_NOT_ALLOWED              HttpMessage = 11
	HttpMessage_REQUEST_TIMEOUT                 HttpMessage = 12
	HttpMessage_TOO_MANY_REQUESTS               HttpMessage = 13
	HttpMessage_CLIENT_CLOSED_REQUEST           HttpMessage = 14
	HttpMessage_INTERNAL_SERVER_ERROR           HttpMessage = 15
	HttpMessage_NOT_IMPLEMENTED                 HttpMessage = 16
	HttpMessage_BAD_GATEWAY                     HttpMessage = 17
	HttpMessage_SERVICE_UNAVAILABLE             HttpMessage = 18
	HttpMessage_NETWORK_AUTHENTICATION_REQUIRED HttpMessage = 19
	HttpMessage_NETWORK_CONNECT_TIMEOUT_ERROR   HttpMessage = 20
)

// Enum value maps for HttpMessage.
var (
	HttpMessage_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "PROCESSING",
		2:  "OK",
		3:  "CREATED",
		4:  "NO_CONTENT",
		5:  "PARTIAL_CONTENT",
		6:  "FOUND",
		7:  "BAD_REQUEST",
		8:  "UNAUTHORIZED",
		9:  "FORBIDDEN",
		10: "NOT_FOUND",
		11: "METHOD_NOT_ALLOWED",
		12: "REQUEST_TIMEOUT",
		13: "TOO_MANY_REQUESTS",
		14: "CLIENT_CLOSED_REQUEST",
		15: "INTERNAL_SERVER_ERROR",
		16: "NOT_IMPLEMENTED",
		17: "BAD_GATEWAY",
		18: "SERVICE_UNAVAILABLE",
		19: "NETWORK_AUTHENTICATION_REQUIRED",
		20: "NETWORK_CONNECT_TIMEOUT_ERROR",
	}
	HttpMessage_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"PROCESSING":                      1,
		"OK":                              2,
		"CREATED":                         3,
		"NO_CONTENT":                      4,
		"PARTIAL_CONTENT":                 5,
		"FOUND":                           6,
		"BAD_REQUEST":                     7,
		"UNAUTHORIZED":                    8,
		"FORBIDDEN":                       9,
		"NOT_FOUND":                       10,
		"METHOD_NOT_ALLOWED":              11,
		"REQUEST_TIMEOUT":                 12,
		"TOO_MANY_REQUESTS":               13,
		"CLIENT_CLOSED_REQUEST":           14,
		"INTERNAL_SERVER_ERROR":           15,
		"NOT_IMPLEMENTED":                 16,
		"BAD_GATEWAY":                     17,
		"SERVICE_UNAVAILABLE":             18,
		"NETWORK_AUTHENTICATION_REQUIRED": 19,
		"NETWORK_CONNECT_TIMEOUT_ERROR":   20,
	}
)

func (x HttpMessage) Enum() *HttpMessage {
	p := new(HttpMessage)
	*p = x
	return p
}

func (x HttpMessage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpMessage) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_go_grpc_redis_proto_enumTypes[1].Descriptor()
}

func (HttpMessage) Type() protoreflect.EnumType {
	return &file_proto_go_grpc_redis_proto_enumTypes[1]
}

func (x HttpMessage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpMessage.Descriptor instead.
func (HttpMessage) EnumDescriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{1}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode    int32       `protobuf:"varint,1,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	ResultMessage HttpMessage `protobuf:"varint,2,opt,name=resultMessage,proto3,enum=goredis.HttpMessage" json:"resultMessage,omitempty"`
	Result        string      `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *GetResponse) GetResultMessage() HttpMessage {
	if x != nil {
		return x.ResultMessage
	}
	return HttpMessage_UNSPECIFIED
}

func (x *GetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{2}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode    int32       `protobuf:"varint,1,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	ResultMessage HttpMessage `protobuf:"varint,2,opt,name=resultMessage,proto3,enum=goredis.HttpMessage" json:"resultMessage,omitempty"`
	Result        string      `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{3}
}

func (x *SetResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *SetResponse) GetResultMessage() HttpMessage {
	if x != nil {
		return x.ResultMessage
	}
	return HttpMessage_UNSPECIFIED
}

func (x *SetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode    int32       `protobuf:"varint,1,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	ResultMessage HttpMessage `protobuf:"varint,2,opt,name=resultMessage,proto3,enum=goredis.HttpMessage" json:"resultMessage,omitempty"`
	Result        string      `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *DeleteResponse) GetResultMessage() HttpMessage {
	if x != nil {
		return x.ResultMessage
	}
	return HttpMessage_UNSPECIFIED
}

func (x *DeleteResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode    int32       `protobuf:"varint,1,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	ResultMessage HttpMessage `protobuf:"varint,2,opt,name=resultMessage,proto3,enum=goredis.HttpMessage" json:"resultMessage,omitempty"`
	Result        string      `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *UpdateResponse) GetResultMessage() HttpMessage {
	if x != nil {
		return x.ResultMessage
	}
	return HttpMessage_UNSPECIFIED
}

func (x *UpdateResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{8}
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message uint32 `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_go_grpc_redis_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_go_grpc_redis_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_proto_go_grpc_redis_proto_rawDescGZIP(), []int{9}
}

func (x *HealthResponse) GetMessage() uint32 {
	if x != nil {
		return x.Message
	}
	return 0
}

var File_proto_go_grpc_redis_proto protoreflect.FileDescriptor

var file_proto_go_grpc_redis_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x6f, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x81,
	0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x37, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04,
	0x2a, 0xb5, 0x03, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52,
	0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x0b, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0x0d, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x47, 0x41,
	0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x12,
	0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x13, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x14, 0x32, 0xe0, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x13, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x53,
	0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x0e, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x67, 0x6f, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_go_grpc_redis_proto_rawDescOnce sync.Once
	file_proto_go_grpc_redis_proto_rawDescData = file_proto_go_grpc_redis_proto_rawDesc
)

func file_proto_go_grpc_redis_proto_rawDescGZIP() []byte {
	file_proto_go_grpc_redis_proto_rawDescOnce.Do(func() {
		file_proto_go_grpc_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_go_grpc_redis_proto_rawDescData)
	})
	return file_proto_go_grpc_redis_proto_rawDescData
}

var file_proto_go_grpc_redis_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_go_grpc_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_go_grpc_redis_proto_goTypes = []interface{}{
	(ResultStatus)(0),      // 0: goredis.ResultStatus
	(HttpMessage)(0),       // 1: goredis.HttpMessage
	(*GetRequest)(nil),     // 2: goredis.GetRequest
	(*GetResponse)(nil),    // 3: goredis.GetResponse
	(*SetRequest)(nil),     // 4: goredis.SetRequest
	(*SetResponse)(nil),    // 5: goredis.SetResponse
	(*DeleteRequest)(nil),  // 6: goredis.DeleteRequest
	(*DeleteResponse)(nil), // 7: goredis.DeleteResponse
	(*UpdateRequest)(nil),  // 8: goredis.UpdateRequest
	(*UpdateResponse)(nil), // 9: goredis.UpdateResponse
	(*Empty)(nil),          // 10: goredis.Empty
	(*HealthResponse)(nil), // 11: goredis.HealthResponse
}
var file_proto_go_grpc_redis_proto_depIdxs = []int32{
	1,  // 0: goredis.GetResponse.resultMessage:type_name -> goredis.HttpMessage
	1,  // 1: goredis.SetResponse.resultMessage:type_name -> goredis.HttpMessage
	1,  // 2: goredis.DeleteResponse.resultMessage:type_name -> goredis.HttpMessage
	1,  // 3: goredis.UpdateResponse.resultMessage:type_name -> goredis.HttpMessage
	2,  // 4: goredis.CacheService.Get:input_type -> goredis.GetRequest
	4,  // 5: goredis.CacheService.Set:input_type -> goredis.SetRequest
	6,  // 6: goredis.CacheService.Delete:input_type -> goredis.DeleteRequest
	10, // 7: goredis.CacheService.Health:input_type -> goredis.Empty
	3,  // 8: goredis.CacheService.Get:output_type -> goredis.GetResponse
	5,  // 9: goredis.CacheService.Set:output_type -> goredis.SetResponse
	7,  // 10: goredis.CacheService.Delete:output_type -> goredis.DeleteResponse
	11, // 11: goredis.CacheService.Health:output_type -> goredis.HealthResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_go_grpc_redis_proto_init() }
func file_proto_go_grpc_redis_proto_init() {
	if File_proto_go_grpc_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_go_grpc_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_go_grpc_redis_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_proto_go_grpc_redis_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_go_grpc_redis_proto_goTypes,
		DependencyIndexes: file_proto_go_grpc_redis_proto_depIdxs,
		EnumInfos:         file_proto_go_grpc_redis_proto_enumTypes,
		MessageInfos:      file_proto_go_grpc_redis_proto_msgTypes,
	}.Build()
	File_proto_go_grpc_redis_proto = out.File
	file_proto_go_grpc_redis_proto_rawDesc = nil
	file_proto_go_grpc_redis_proto_goTypes = nil
	file_proto_go_grpc_redis_proto_depIdxs = nil
}
