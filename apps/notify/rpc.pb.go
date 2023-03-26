// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/notify/pb/rpc.proto

package notify

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

type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信模版的Id
	// @gotags: bson:"template_id" json:"template_id"
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id" bson:"template_id"`
	// 模版参数
	// @gotags: bson:"template_params" json:"template_params"
	TemplateParams []string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params" bson:"template_params"`
	// 用户
	// @gotags: bson:"users" json:"users"
	Users []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users" bson:"users"`
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *SendSMSRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendSMSRequest) GetTemplateParams() []string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

func (x *SendSMSRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

type SendSmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送成功的电话列表
	// @gotags: json:"successed_numbers"
	SuccessedNumbers []string `protobuf:"bytes,1,rep,name=successed_numbers,json=successedNumbers,proto3" json:"successed_numbers"`
}

func (x *SendSmsResponse) Reset() {
	*x = SendSmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsResponse) ProtoMessage() {}

func (x *SendSmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsResponse.ProtoReflect.Descriptor instead.
func (*SendSmsResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *SendSmsResponse) GetSuccessedNumbers() []string {
	if x != nil {
		return x.SuccessedNumbers
	}
	return nil
}

type SendVoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信模版的Id
	// @gotags: bson:"template_id" json:"template_id"
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id" bson:"template_id"`
	// 模版参数
	// @gotags: bson:"template_params" json:"template_params"
	TemplateParams []string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params" bson:"template_params"`
	// 用户
	// @gotags: bson:"users" json:"users"
	Users []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users" bson:"users"`
	// 播放次数，可选，最多3次，默认2次
	// @gotags: bson:"play_times" json:"play_times"
	PlayTimes uint64 `protobuf:"varint,4,opt,name=play_times,json=playTimes,proto3" json:"play_times" bson:"play_times"`
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	// @gotags: bson:"session_context" json:"session_context"
	SessionContext string `protobuf:"bytes,5,opt,name=session_context,json=sessionContext,proto3" json:"session_context" bson:"session_context"`
}

func (x *SendVoiceRequest) Reset() {
	*x = SendVoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVoiceRequest) ProtoMessage() {}

func (x *SendVoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVoiceRequest.ProtoReflect.Descriptor instead.
func (*SendVoiceRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *SendVoiceRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendVoiceRequest) GetTemplateParams() []string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

func (x *SendVoiceRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SendVoiceRequest) GetPlayTimes() uint64 {
	if x != nil {
		return x.PlayTimes
	}
	return 0
}

func (x *SendVoiceRequest) GetSessionContext() string {
	if x != nil {
		return x.SessionContext
	}
	return ""
}

type SendVoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 呼叫Id
	// @gotags: bson:"call_id" json:"call_id"
	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id" bson:"call_id"`
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	// @gotags: bson:"session_context" json:"session_context"
	SessionContext string `protobuf:"bytes,2,opt,name=session_context,json=sessionContext,proto3" json:"session_context" bson:"session_context"`
}

func (x *SendVoiceResponse) Reset() {
	*x = SendVoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVoiceResponse) ProtoMessage() {}

func (x *SendVoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVoiceResponse.ProtoReflect.Descriptor instead.
func (*SendVoiceResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *SendVoiceResponse) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *SendVoiceResponse) GetSessionContext() string {
	if x != nil {
		return x.SessionContext
	}
	return ""
}

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	// @gotags: json:"users" bson:"users" validate:"required"
	Users []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users" bson:"users" validate:"required"`
	// 告警时标题
	// @gotags: json:"title" bson:"title" validate:"required"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" bson:"title" validate:"required"`
	// 告警内容
	// @gotags: json:"content" bson:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" bson:"content"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *SendMailRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SendMailRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendMailRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送成功的邮件列表
	// @gotags: json:"successed_mails"
	SuccessedMails []string `protobuf:"bytes,1,rep,name=successed_mails,json=successedMails,proto3" json:"successed_mails"`
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *SendMailResponse) GetSuccessedMails() []string {
	if x != nil {
		return x.SuccessedMails
	}
	return nil
}

type SendIMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	// @gotags: json:"users" bson:"users" validate:"required"
	Users []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users" bson:"users" validate:"required"`
	// 告警时标题
	// @gotags: json:"title" bson:"title" validate:"required"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" bson:"title" validate:"required"`
	// 告警内容
	// @gotags: json:"content" bson:"content"
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content" bson:"content"`
}

func (x *SendIMRequest) Reset() {
	*x = SendIMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendIMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendIMRequest) ProtoMessage() {}

func (x *SendIMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendIMRequest.ProtoReflect.Descriptor instead.
func (*SendIMRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *SendIMRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SendIMRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendIMRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendImResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendImResponse) Reset() {
	*x = SendImResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendImResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendImResponse) ProtoMessage() {}

func (x *SendImResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendImResponse.ProtoReflect.Descriptor instead.
func (*SendImResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_rpc_proto_rawDescGZIP(), []int{7}
}

var File_apps_notify_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_notify_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x70, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3e, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x55, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x55, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64,
	0x49, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x93, 0x03, 0x0a, 0x03, 0x52,
	0x50, 0x43, 0x12, 0x60, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x12, 0x29, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d,
	0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x4d, 0x12, 0x28, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x4d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_notify_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_notify_pb_rpc_proto_rawDescData = file_apps_notify_pb_rpc_proto_rawDesc
)

func file_apps_notify_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_notify_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_notify_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_notify_pb_rpc_proto_rawDescData)
	})
	return file_apps_notify_pb_rpc_proto_rawDescData
}

var file_apps_notify_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apps_notify_pb_rpc_proto_goTypes = []interface{}{
	(*SendSMSRequest)(nil),    // 0: infraboard.mcenter.notify.SendSMSRequest
	(*SendSmsResponse)(nil),   // 1: infraboard.mcenter.notify.SendSmsResponse
	(*SendVoiceRequest)(nil),  // 2: infraboard.mcenter.notify.SendVoiceRequest
	(*SendVoiceResponse)(nil), // 3: infraboard.mcenter.notify.SendVoiceResponse
	(*SendMailRequest)(nil),   // 4: infraboard.mcenter.notify.SendMailRequest
	(*SendMailResponse)(nil),  // 5: infraboard.mcenter.notify.SendMailResponse
	(*SendIMRequest)(nil),     // 6: infraboard.mcenter.notify.SendIMRequest
	(*SendImResponse)(nil),    // 7: infraboard.mcenter.notify.SendImResponse
}
var file_apps_notify_pb_rpc_proto_depIdxs = []int32{
	0, // 0: infraboard.mcenter.notify.RPC.SendSMS:input_type -> infraboard.mcenter.notify.SendSMSRequest
	2, // 1: infraboard.mcenter.notify.RPC.SendVoice:input_type -> infraboard.mcenter.notify.SendVoiceRequest
	4, // 2: infraboard.mcenter.notify.RPC.SendMail:input_type -> infraboard.mcenter.notify.SendMailRequest
	6, // 3: infraboard.mcenter.notify.RPC.SendIM:input_type -> infraboard.mcenter.notify.SendIMRequest
	1, // 4: infraboard.mcenter.notify.RPC.SendSMS:output_type -> infraboard.mcenter.notify.SendSmsResponse
	3, // 5: infraboard.mcenter.notify.RPC.SendVoice:output_type -> infraboard.mcenter.notify.SendVoiceResponse
	5, // 6: infraboard.mcenter.notify.RPC.SendMail:output_type -> infraboard.mcenter.notify.SendMailResponse
	7, // 7: infraboard.mcenter.notify.RPC.SendIM:output_type -> infraboard.mcenter.notify.SendImResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_notify_pb_rpc_proto_init() }
func file_apps_notify_pb_rpc_proto_init() {
	if File_apps_notify_pb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_notify_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsResponse); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVoiceRequest); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVoiceResponse); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResponse); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendIMRequest); i {
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
		file_apps_notify_pb_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendImResponse); i {
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
			RawDescriptor: file_apps_notify_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_notify_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_notify_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_notify_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_notify_pb_rpc_proto = out.File
	file_apps_notify_pb_rpc_proto_rawDesc = nil
	file_apps_notify_pb_rpc_proto_goTypes = nil
	file_apps_notify_pb_rpc_proto_depIdxs = nil
}
