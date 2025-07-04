// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: config.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListConfigKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *ListConfigKeysRequest) Reset() {
	*x = ListConfigKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigKeysRequest) ProtoMessage() {}

func (x *ListConfigKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigKeysRequest.ProtoReflect.Descriptor instead.
func (*ListConfigKeysRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *ListConfigKeysRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type ListConfigKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configkeys []string `protobuf:"bytes,1,rep,name=configkeys,proto3" json:"configkeys,omitempty"`
}

func (x *ListConfigKeysResponse) Reset() {
	*x = ListConfigKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigKeysResponse) ProtoMessage() {}

func (x *ListConfigKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigKeysResponse.ProtoReflect.Descriptor instead.
func (*ListConfigKeysResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *ListConfigKeysResponse) GetConfigkeys() []string {
	if x != nil {
		return x.Configkeys
	}
	return nil
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configkey string `protobuf:"bytes,1,opt,name=configkey,proto3" json:"configkey,omitempty"`
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *GetConfigRequest) GetConfigkey() string {
	if x != nil {
		return x.Configkey
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigKey string `protobuf:"bytes,1,opt,name=config_key,json=configKey,proto3" json:"config_key,omitempty"`
	Schema    string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Ui        string `protobuf:"bytes,3,opt,name=ui,proto3" json:"ui,omitempty"`
	Settings  string `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigResponse) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *ConfigResponse) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *ConfigResponse) GetUi() string {
	if x != nil {
		return x.Ui
	}
	return ""
}

func (x *ConfigResponse) GetSettings() string {
	if x != nil {
		return x.Settings
	}
	return ""
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configkey string `protobuf:"bytes,1,opt,name=configkey,proto3" json:"configkey,omitempty"`
	Settings  string `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateConfigRequest) GetConfigkey() string {
	if x != nil {
		return x.Configkey
	}
	return ""
}

func (x *UpdateConfigRequest) GetSettings() string {
	if x != nil {
		return x.Settings
	}
	return ""
}

type UpdateConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigKey string `protobuf:"bytes,1,opt,name=config_key,json=configKey,proto3" json:"config_key,omitempty"`
	Settings  string `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateConfigResponse) Reset() {
	*x = UpdateConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigResponse) ProtoMessage() {}

func (x *UpdateConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateConfigResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConfigResponse) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *UpdateConfigResponse) GetSettings() string {
	if x != nil {
		return x.Settings
	}
	return ""
}

type GetConfigValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configkey string `protobuf:"bytes,1,opt,name=configkey,proto3" json:"configkey,omitempty"`
	Selector  string `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *GetConfigValueRequest) Reset() {
	*x = GetConfigValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigValueRequest) ProtoMessage() {}

func (x *GetConfigValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigValueRequest.ProtoReflect.Descriptor instead.
func (*GetConfigValueRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

func (x *GetConfigValueRequest) GetConfigkey() string {
	if x != nil {
		return x.Configkey
	}
	return ""
}

func (x *GetConfigValueRequest) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

type GetConfigValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigKey string     `protobuf:"bytes,1,opt,name=config_key,json=configKey,proto3" json:"config_key,omitempty"`
	Selector  string     `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
	Value     *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetConfigValueResponse) Reset() {
	*x = GetConfigValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigValueResponse) ProtoMessage() {}

func (x *GetConfigValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigValueResponse.ProtoReflect.Descriptor instead.
func (*GetConfigValueResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{7}
}

func (x *GetConfigValueResponse) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *GetConfigValueResponse) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *GetConfigValueResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetConfigValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configkey string     `protobuf:"bytes,1,opt,name=configkey,proto3" json:"configkey,omitempty"`
	Selector  string     `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
	Value     *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetConfigValueRequest) Reset() {
	*x = SetConfigValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigValueRequest) ProtoMessage() {}

func (x *SetConfigValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigValueRequest.ProtoReflect.Descriptor instead.
func (*SetConfigValueRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{8}
}

func (x *SetConfigValueRequest) GetConfigkey() string {
	if x != nil {
		return x.Configkey
	}
	return ""
}

func (x *SetConfigValueRequest) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *SetConfigValueRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetConfigValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigKey string     `protobuf:"bytes,1,opt,name=config_key,json=configKey,proto3" json:"config_key,omitempty"`
	Selector  string     `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
	OldValue  *anypb.Any `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
}

func (x *SetConfigValueResponse) Reset() {
	*x = SetConfigValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigValueResponse) ProtoMessage() {}

func (x *SetConfigValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigValueResponse.ProtoReflect.Descriptor instead.
func (*SetConfigValueResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{9}
}

func (x *SetConfigValueResponse) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *SetConfigValueResponse) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *SetConfigValueResponse) GetOldValue() *anypb.Any {
	if x != nil {
		return x.OldValue
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x22, 0x38, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x22, 0x73, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x75, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x4f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x51, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x51, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x7f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7d, 0x0a, 0x15, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x31,
	0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xe3, 0x07, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xb8, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x41, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x24, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x20, 0x6b, 0x65, 0x79, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x7d, 0x12, 0xb1,
	0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72,
	0x92, 0x41, 0x4d, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x30,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x20, 0x73, 0x63,
	0x68, 0x61, 0x6d, 0x61, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65,
	0x79, 0x7d, 0x12, 0xb8, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d,
	0x92, 0x41, 0x45, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x20, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01,
	0x2a, 0x1a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0xd4, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x82, 0x01, 0x92, 0x41, 0x4e, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x20, 0x4b, 0x65, 0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x1a, 0x20, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x20, 0x6b, 0x65, 0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6b,
	0x65, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x6b, 0x65, 0x79, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x7b, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x7d, 0x12, 0xd0, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x92, 0x41, 0x4b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x53, 0x65, 0x74, 0x20,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x20, 0x4b, 0x65, 0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1d, 0x53, 0x65, 0x74, 0x20, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x1a, 0x29, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x6b, 0x65, 0x79, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x7b, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x7d, 0x42, 0x77, 0x92, 0x41, 0x4b, 0x12, 0x05, 0x32, 0x03,
	0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x72, 0x3e, 0x0a, 0x17, 0x4d, 0x58, 0x20, 0x61, 0x64,
	0x73, 0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x23, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6e, 0x68, 0x75, 0x61, 0x78, 0x69, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6e, 0x68, 0x75, 0x61, 0x78, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_config_proto_goTypes = []interface{}{
	(*ListConfigKeysRequest)(nil),  // 0: config.ListConfigKeysRequest
	(*ListConfigKeysResponse)(nil), // 1: config.ListConfigKeysResponse
	(*GetConfigRequest)(nil),       // 2: config.GetConfigRequest
	(*ConfigResponse)(nil),         // 3: config.ConfigResponse
	(*UpdateConfigRequest)(nil),    // 4: config.UpdateConfigRequest
	(*UpdateConfigResponse)(nil),   // 5: config.UpdateConfigResponse
	(*GetConfigValueRequest)(nil),  // 6: config.GetConfigValueRequest
	(*GetConfigValueResponse)(nil), // 7: config.GetConfigValueResponse
	(*SetConfigValueRequest)(nil),  // 8: config.SetConfigValueRequest
	(*SetConfigValueResponse)(nil), // 9: config.SetConfigValueResponse
	(*anypb.Any)(nil),              // 10: google.protobuf.Any
}
var file_config_proto_depIdxs = []int32{
	10, // 0: config.GetConfigValueResponse.value:type_name -> google.protobuf.Any
	10, // 1: config.SetConfigValueRequest.value:type_name -> google.protobuf.Any
	10, // 2: config.SetConfigValueResponse.old_value:type_name -> google.protobuf.Any
	0,  // 3: config.ConfigService.ListConfigKeys:input_type -> config.ListConfigKeysRequest
	2,  // 4: config.ConfigService.GetConfig:input_type -> config.GetConfigRequest
	4,  // 5: config.ConfigService.UpdateConfig:input_type -> config.UpdateConfigRequest
	6,  // 6: config.ConfigService.GetConfigValue:input_type -> config.GetConfigValueRequest
	8,  // 7: config.ConfigService.SetConfigValue:input_type -> config.SetConfigValueRequest
	1,  // 8: config.ConfigService.ListConfigKeys:output_type -> config.ListConfigKeysResponse
	3,  // 9: config.ConfigService.GetConfig:output_type -> config.ConfigResponse
	5,  // 10: config.ConfigService.UpdateConfig:output_type -> config.UpdateConfigResponse
	7,  // 11: config.ConfigService.GetConfigValue:output_type -> config.GetConfigValueResponse
	9,  // 12: config.ConfigService.SetConfigValue:output_type -> config.SetConfigValueResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigKeysRequest); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigKeysResponse); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigResponse); i {
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
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigValueRequest); i {
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
		file_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigValueResponse); i {
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
		file_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigValueRequest); i {
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
		file_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigValueResponse); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
