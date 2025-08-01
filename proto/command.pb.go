// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: command.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	common "github.com/hnhuaxi/proto/proto/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImportScope int32

const (
	ImportScope_IMPORT_ADS             ImportScope = 0
	ImportScope_IMPORT_CAMPAIGNS       ImportScope = 1
	ImportScope_IMPORT_ADGROUPS        ImportScope = 2
	ImportScope_IMPORT_USER_ACTION_SET ImportScope = 3
	ImportScope_IMPORT_ADCREATIVES     ImportScope = 4
	ImportScope_IMPORT_IMAGES          ImportScope = 5
)

// Enum value maps for ImportScope.
var (
	ImportScope_name = map[int32]string{
		0: "IMPORT_ADS",
		1: "IMPORT_CAMPAIGNS",
		2: "IMPORT_ADGROUPS",
		3: "IMPORT_USER_ACTION_SET",
		4: "IMPORT_ADCREATIVES",
		5: "IMPORT_IMAGES",
	}
	ImportScope_value = map[string]int32{
		"IMPORT_ADS":             0,
		"IMPORT_CAMPAIGNS":       1,
		"IMPORT_ADGROUPS":        2,
		"IMPORT_USER_ACTION_SET": 3,
		"IMPORT_ADCREATIVES":     4,
		"IMPORT_IMAGES":          5,
	}
)

func (x ImportScope) Enum() *ImportScope {
	p := new(ImportScope)
	*p = x
	return p
}

func (x ImportScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportScope) Descriptor() protoreflect.EnumDescriptor {
	return file_command_proto_enumTypes[0].Descriptor()
}

func (ImportScope) Type() protoreflect.EnumType {
	return &file_command_proto_enumTypes[0]
}

func (x ImportScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportScope.Descriptor instead.
func (ImportScope) EnumDescriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

// ImportResourceCommand 导入资源
type ImportResourceCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string           `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId string           `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Scopes    []ImportScope    `protobuf:"varint,3,rep,packed,name=scopes,proto3,enum=account.ImportScope" json:"scopes,omitempty"`
	Filters   []*common.Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ImportResourceCommand) Reset() {
	*x = ImportResourceCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResourceCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResourceCommand) ProtoMessage() {}

func (x *ImportResourceCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResourceCommand.ProtoReflect.Descriptor instead.
func (*ImportResourceCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *ImportResourceCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ImportResourceCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ImportResourceCommand) GetScopes() []ImportScope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ImportResourceCommand) GetFilters() []*common.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ReceivedResourcesEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ReceivedResourcesEvent) Reset() {
	*x = ReceivedResourcesEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedResourcesEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedResourcesEvent) ProtoMessage() {}

func (x *ReceivedResourcesEvent) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedResourcesEvent.ProtoReflect.Descriptor instead.
func (*ReceivedResourcesEvent) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (x *ReceivedResourcesEvent) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// 提交过审
type AdCommitReviewCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider     string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId    string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AdcreativeId string `protobuf:"bytes,3,opt,name=adcreative_id,json=adcreativeId,proto3" json:"adcreative_id,omitempty"`
}

func (x *AdCommitReviewCommand) Reset() {
	*x = AdCommitReviewCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdCommitReviewCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdCommitReviewCommand) ProtoMessage() {}

func (x *AdCommitReviewCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdCommitReviewCommand.ProtoReflect.Descriptor instead.
func (*AdCommitReviewCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *AdCommitReviewCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AdCommitReviewCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AdCommitReviewCommand) GetAdcreativeId() string {
	if x != nil {
		return x.AdcreativeId
	}
	return ""
}

// StartRunning 开始广告
type AdStartCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *AdStartCommand) Reset() {
	*x = AdStartCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdStartCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdStartCommand) ProtoMessage() {}

func (x *AdStartCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdStartCommand.ProtoReflect.Descriptor instead.
func (*AdStartCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

func (x *AdStartCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AdStartCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

// Pause 暂停广告
type AdPauseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *AdPauseCommand) Reset() {
	*x = AdPauseCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdPauseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdPauseCommand) ProtoMessage() {}

func (x *AdPauseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdPauseCommand.ProtoReflect.Descriptor instead.
func (*AdPauseCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{4}
}

func (x *AdPauseCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AdPauseCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

// Stop 停止广告
type AdStopCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *AdStopCommand) Reset() {
	*x = AdStopCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdStopCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdStopCommand) ProtoMessage() {}

func (x *AdStopCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdStopCommand.ProtoReflect.Descriptor instead.
func (*AdStopCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{5}
}

func (x *AdStopCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AdStopCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

// 改变出价
type AdChangePriceCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *AdChangePriceCommand) Reset() {
	*x = AdChangePriceCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdChangePriceCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdChangePriceCommand) ProtoMessage() {}

func (x *AdChangePriceCommand) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdChangePriceCommand.ProtoReflect.Descriptor instead.
func (*AdChangePriceCommand) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{6}
}

func (x *AdChangePriceCommand) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *AdChangePriceCommand) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa,
	0x01, 0x0a, 0x15, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x15, 0x41, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x64, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0e,
	0x41, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0e, 0x41, 0x64, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0d, 0x41, 0x64, 0x53, 0x74, 0x6f, 0x70,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x51, 0x0a, 0x14, 0x41, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0x8f, 0x01, 0x0a, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x41, 0x44, 0x53, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x41, 0x44, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x53, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x53, 0x10, 0x05, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6e, 0x68, 0x75, 0x61, 0x78, 0x69, 0x2f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_command_proto_goTypes = []interface{}{
	(ImportScope)(0),               // 0: account.ImportScope
	(*ImportResourceCommand)(nil),  // 1: account.ImportResourceCommand
	(*ReceivedResourcesEvent)(nil), // 2: account.ReceivedResourcesEvent
	(*AdCommitReviewCommand)(nil),  // 3: account.AdCommitReviewCommand
	(*AdStartCommand)(nil),         // 4: account.AdStartCommand
	(*AdPauseCommand)(nil),         // 5: account.AdPauseCommand
	(*AdStopCommand)(nil),          // 6: account.AdStopCommand
	(*AdChangePriceCommand)(nil),   // 7: account.AdChangePriceCommand
	(*common.Filter)(nil),          // 8: common.Filter
	(*Account)(nil),                // 9: account.Account
}
var file_command_proto_depIdxs = []int32{
	0, // 0: account.ImportResourceCommand.scopes:type_name -> account.ImportScope
	8, // 1: account.ImportResourceCommand.filters:type_name -> common.Filter
	9, // 2: account.ReceivedResourcesEvent.accounts:type_name -> account.Account
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	file_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResourceCommand); i {
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
		file_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedResourcesEvent); i {
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
		file_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdCommitReviewCommand); i {
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
		file_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdStartCommand); i {
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
		file_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdPauseCommand); i {
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
		file_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdStopCommand); i {
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
		file_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdChangePriceCommand); i {
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
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		EnumInfos:         file_command_proto_enumTypes,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
