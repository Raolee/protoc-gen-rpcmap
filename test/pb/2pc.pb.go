// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: 2pc.proto

package service

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

type Result int32

const (
	Result_SUCCESS Result = 0
	Result_ALREADY Result = 1
	Result_FAILURE Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "SUCCESS",
		1: "ALREADY",
		2: "FAILURE",
	}
	Result_value = map[string]int32{
		"SUCCESS": 0,
		"ALREADY": 1,
		"FAILURE": 2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file__2pc_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file__2pc_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{0}
}

type Transaction_Status int32

const (
	Transaction_PREPARED     Transaction_Status = 0
	Transaction_COMMITTING   Transaction_Status = 10
	Transaction_COMMITTED    Transaction_Status = 11
	Transaction_ROLLING_BACK Transaction_Status = 20
	Transaction_ROLLED_BACK  Transaction_Status = 21
)

// Enum value maps for Transaction_Status.
var (
	Transaction_Status_name = map[int32]string{
		0:  "PREPARED",
		10: "COMMITTING",
		11: "COMMITTED",
		20: "ROLLING_BACK",
		21: "ROLLED_BACK",
	}
	Transaction_Status_value = map[string]int32{
		"PREPARED":     0,
		"COMMITTING":   10,
		"COMMITTED":    11,
		"ROLLING_BACK": 20,
		"ROLLED_BACK":  21,
	}
)

func (x Transaction_Status) Enum() *Transaction_Status {
	p := new(Transaction_Status)
	*p = x
	return p
}

func (x Transaction_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transaction_Status) Descriptor() protoreflect.EnumDescriptor {
	return file__2pc_proto_enumTypes[1].Descriptor()
}

func (Transaction_Status) Type() protoreflect.EnumType {
	return &file__2pc_proto_enumTypes[1]
}

func (x Transaction_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transaction_Status.Descriptor instead.
func (Transaction_Status) EnumDescriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{0, 0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId   string             `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Domain string             `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Status Transaction_Status `protobuf:"varint,10,opt,name=status,proto3,enum=twophasecommit.Transaction_Status" json:"status,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Transaction) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Transaction) GetStatus() Transaction_Status {
	if x != nil {
		return x.Status
	}
	return Transaction_PREPARED
}

type PrepareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId   string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *PrepareRequest) Reset() {
	*x = PrepareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRequest) ProtoMessage() {}

func (x *PrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRequest.ProtoReflect.Descriptor instead.
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{1}
}

func (x *PrepareRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *PrepareRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type PrepareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
}

func (x *PrepareResponse) Reset() {
	*x = PrepareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareResponse) ProtoMessage() {}

func (x *PrepareResponse) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareResponse.ProtoReflect.Descriptor instead.
func (*PrepareResponse) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareResponse) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type CommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId   string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *CommitRequest) Reset() {
	*x = CommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRequest) ProtoMessage() {}

func (x *CommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRequest.ProtoReflect.Descriptor instead.
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{3}
}

func (x *CommitRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *CommitRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type CommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=twophasecommit.Result" json:"result,omitempty"`
}

func (x *CommitResponse) Reset() {
	*x = CommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResponse) ProtoMessage() {}

func (x *CommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResponse.ProtoReflect.Descriptor instead.
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{4}
}

func (x *CommitResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_SUCCESS
}

type RollbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId   string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *RollbackRequest) Reset() {
	*x = RollbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackRequest) ProtoMessage() {}

func (x *RollbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackRequest.ProtoReflect.Descriptor instead.
func (*RollbackRequest) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{5}
}

func (x *RollbackRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *RollbackRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type RollbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=twophasecommit.Result" json:"result,omitempty"`
}

func (x *RollbackResponse) Reset() {
	*x = RollbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackResponse) ProtoMessage() {}

func (x *RollbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackResponse.ProtoReflect.Descriptor instead.
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{6}
}

func (x *RollbackResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_SUCCESS
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId    string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Success bool   `protobuf:"varint,11,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{7}
}

func (x *ReportRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *ReportRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ReportRequest) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=twophasecommit.Result" json:"result,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file__2pc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file__2pc_proto_rawDescGZIP(), []int{8}
}

func (x *ReportResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_SUCCESS
}

var File__2pc_proto protoreflect.FileDescriptor

var file__2pc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x32, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x77, 0x6f,
	0x70, 0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x74,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x4f, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x14, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x15, 0x22, 0x3d,
	0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x26, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x42, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x6f, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x56, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x40, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2a, 0x2f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x02, 0x32, 0xcc, 0x02, 0x0a, 0x15, 0x54, 0x77, 0x6f, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x77, 0x6f, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x77, 0x6f, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x12, 0x1f, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x77, 0x6f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x32, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__2pc_proto_rawDescOnce sync.Once
	file__2pc_proto_rawDescData = file__2pc_proto_rawDesc
)

func file__2pc_proto_rawDescGZIP() []byte {
	file__2pc_proto_rawDescOnce.Do(func() {
		file__2pc_proto_rawDescData = protoimpl.X.CompressGZIP(file__2pc_proto_rawDescData)
	})
	return file__2pc_proto_rawDescData
}

var file__2pc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file__2pc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file__2pc_proto_goTypes = []interface{}{
	(Result)(0),              // 0: twophasecommit.Result
	(Transaction_Status)(0),  // 1: twophasecommit.Transaction.Status
	(*Transaction)(nil),      // 2: twophasecommit.Transaction
	(*PrepareRequest)(nil),   // 3: twophasecommit.PrepareRequest
	(*PrepareResponse)(nil),  // 4: twophasecommit.PrepareResponse
	(*CommitRequest)(nil),    // 5: twophasecommit.CommitRequest
	(*CommitResponse)(nil),   // 6: twophasecommit.CommitResponse
	(*RollbackRequest)(nil),  // 7: twophasecommit.RollbackRequest
	(*RollbackResponse)(nil), // 8: twophasecommit.RollbackResponse
	(*ReportRequest)(nil),    // 9: twophasecommit.ReportRequest
	(*ReportResponse)(nil),   // 10: twophasecommit.ReportResponse
}
var file__2pc_proto_depIdxs = []int32{
	1,  // 0: twophasecommit.Transaction.status:type_name -> twophasecommit.Transaction.Status
	0,  // 1: twophasecommit.CommitResponse.result:type_name -> twophasecommit.Result
	0,  // 2: twophasecommit.RollbackResponse.result:type_name -> twophasecommit.Result
	0,  // 3: twophasecommit.ReportResponse.result:type_name -> twophasecommit.Result
	3,  // 4: twophasecommit.TwoPhaseCommitService.Prepare:input_type -> twophasecommit.PrepareRequest
	5,  // 5: twophasecommit.TwoPhaseCommitService.Commit:input_type -> twophasecommit.CommitRequest
	7,  // 6: twophasecommit.TwoPhaseCommitService.Rollback:input_type -> twophasecommit.RollbackRequest
	9,  // 7: twophasecommit.TwoPhaseCommitService.Report:input_type -> twophasecommit.ReportRequest
	4,  // 8: twophasecommit.TwoPhaseCommitService.Prepare:output_type -> twophasecommit.PrepareResponse
	6,  // 9: twophasecommit.TwoPhaseCommitService.Commit:output_type -> twophasecommit.CommitResponse
	8,  // 10: twophasecommit.TwoPhaseCommitService.Rollback:output_type -> twophasecommit.RollbackResponse
	10, // 11: twophasecommit.TwoPhaseCommitService.Report:output_type -> twophasecommit.ReportResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file__2pc_proto_init() }
func file__2pc_proto_init() {
	if File__2pc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__2pc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file__2pc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRequest); i {
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
		file__2pc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareResponse); i {
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
		file__2pc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitRequest); i {
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
		file__2pc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResponse); i {
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
		file__2pc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackRequest); i {
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
		file__2pc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackResponse); i {
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
		file__2pc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file__2pc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
			RawDescriptor: file__2pc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__2pc_proto_goTypes,
		DependencyIndexes: file__2pc_proto_depIdxs,
		EnumInfos:         file__2pc_proto_enumTypes,
		MessageInfos:      file__2pc_proto_msgTypes,
	}.Build()
	File__2pc_proto = out.File
	file__2pc_proto_rawDesc = nil
	file__2pc_proto_goTypes = nil
	file__2pc_proto_depIdxs = nil
}
