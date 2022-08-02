// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: apps/code/pb/code.proto

package code

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

type ISSUE_BY int32

const (
	// 通过用户名密码 颁发验证码
	ISSUE_BY_PASSWORD ISSUE_BY = 0
	// 通过AccessToken颁发验证码
	ISSUE_BY_ACCESS_TOKEN ISSUE_BY = 1
)

// Enum value maps for ISSUE_BY.
var (
	ISSUE_BY_name = map[int32]string{
		0: "PASSWORD",
		1: "ACCESS_TOKEN",
	}
	ISSUE_BY_value = map[string]int32{
		"PASSWORD":     0,
		"ACCESS_TOKEN": 1,
	}
)

func (x ISSUE_BY) Enum() *ISSUE_BY {
	p := new(ISSUE_BY)
	*p = x
	return p
}

func (x ISSUE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ISSUE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_code_pb_code_proto_enumTypes[0].Descriptor()
}

func (ISSUE_BY) Type() protoreflect.EnumType {
	return &file_apps_code_pb_code_proto_enumTypes[0]
}

func (x ISSUE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ISSUE_BY.Descriptor instead.
func (ISSUE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_code_pb_code_proto_rawDescGZIP(), []int{0}
}

// Code 验证码
type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 验证码
	// @gotags: bson:"code" json:"code"
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" bson:"code"`
	// 用户Id
	// @gotags: json:"username" validate:"required"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username" validate:"required"`
	// 颁发时间
	// @gotags: bson:"issue_at" json:"issue_at"
	IssueAt int64 `protobuf:"varint,4,opt,name=issue_at,json=issueAt,proto3" json:"issue_at" bson:"issue_at"`
	// 验证码过期时间
	// @gotags: bson:"expired_minite" json:"expired_minite"
	ExpiredMinite uint32 `protobuf:"varint,5,opt,name=expired_minite,json=expiredMinite,proto3" json:"expired_minite" bson:"expired_minite"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_code_pb_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_apps_code_pb_code_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Code) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Code) GetIssueAt() int64 {
	if x != nil {
		return x.IssueAt
	}
	return 0
}

func (x *Code) GetExpiredMinite() uint32 {
	if x != nil {
		return x.ExpiredMinite
	}
	return 0
}

// IssueCodeRequest 验证码申请请求
type IssueCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 颁发方式
	// @gotags: json:"issue_by"
	IssueBy ISSUE_BY `protobuf:"varint,1,opt,name=issue_by,json=issueBy,proto3,enum=infraboard.mcenter.code.ISSUE_BY" json:"issue_by"`
	// 用户名
	// @gotags: json:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	// 密码
	// @gotags: json:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,6,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *IssueCodeRequest) Reset() {
	*x = IssueCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_code_pb_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeRequest) ProtoMessage() {}

func (x *IssueCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCodeRequest.ProtoReflect.Descriptor instead.
func (*IssueCodeRequest) Descriptor() ([]byte, []int) {
	return file_apps_code_pb_code_proto_rawDescGZIP(), []int{1}
}

func (x *IssueCodeRequest) GetIssueBy() ISSUE_BY {
	if x != nil {
		return x.IssueBy
	}
	return ISSUE_BY_PASSWORD
}

func (x *IssueCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueCodeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IssueCodeRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// IssueCodeResponse todo
type IssueCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 颁发后返回的消息, 比如以发送到xxx手机
	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *IssueCodeResponse) Reset() {
	*x = IssueCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_code_pb_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeResponse) ProtoMessage() {}

func (x *IssueCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCodeResponse.ProtoReflect.Descriptor instead.
func (*IssueCodeResponse) Descriptor() ([]byte, []int) {
	return file_apps_code_pb_code_proto_rawDescGZIP(), []int{2}
}

func (x *IssueCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// VerifyCodeRequest 验证码校验请求
type VerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"username" validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" validate:"required"`
	// 验证码
	// @gotags: json:"code" validate:"required"
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" validate:"required"`
}

func (x *VerifyCodeRequest) Reset() {
	*x = VerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_code_pb_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRequest) ProtoMessage() {}

func (x *VerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_apps_code_pb_code_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VerifyCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_apps_code_pb_code_proto protoreflect.FileDescriptor

var file_apps_code_pb_code_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x22, 0xab, 0x01,
	0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49,
	0x53, 0x53, 0x55, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x42, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a,
	0x2a, 0x0a, 0x08, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x0c, 0x0a, 0x08, 0x50,
	0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x32, 0xc2, 0x01, 0x0a, 0x03,
	0x52, 0x50, 0x43, 0x12, 0x62, 0x0a, 0x09, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_code_pb_code_proto_rawDescOnce sync.Once
	file_apps_code_pb_code_proto_rawDescData = file_apps_code_pb_code_proto_rawDesc
)

func file_apps_code_pb_code_proto_rawDescGZIP() []byte {
	file_apps_code_pb_code_proto_rawDescOnce.Do(func() {
		file_apps_code_pb_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_code_pb_code_proto_rawDescData)
	})
	return file_apps_code_pb_code_proto_rawDescData
}

var file_apps_code_pb_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_code_pb_code_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_code_pb_code_proto_goTypes = []interface{}{
	(ISSUE_BY)(0),             // 0: infraboard.mcenter.code.ISSUE_BY
	(*Code)(nil),              // 1: infraboard.mcenter.code.Code
	(*IssueCodeRequest)(nil),  // 2: infraboard.mcenter.code.IssueCodeRequest
	(*IssueCodeResponse)(nil), // 3: infraboard.mcenter.code.IssueCodeResponse
	(*VerifyCodeRequest)(nil), // 4: infraboard.mcenter.code.VerifyCodeRequest
}
var file_apps_code_pb_code_proto_depIdxs = []int32{
	0, // 0: infraboard.mcenter.code.IssueCodeRequest.issue_by:type_name -> infraboard.mcenter.code.ISSUE_BY
	2, // 1: infraboard.mcenter.code.RPC.IssueCode:input_type -> infraboard.mcenter.code.IssueCodeRequest
	4, // 2: infraboard.mcenter.code.RPC.VerifyCode:input_type -> infraboard.mcenter.code.VerifyCodeRequest
	3, // 3: infraboard.mcenter.code.RPC.IssueCode:output_type -> infraboard.mcenter.code.IssueCodeResponse
	1, // 4: infraboard.mcenter.code.RPC.VerifyCode:output_type -> infraboard.mcenter.code.Code
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_code_pb_code_proto_init() }
func file_apps_code_pb_code_proto_init() {
	if File_apps_code_pb_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_code_pb_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_apps_code_pb_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCodeRequest); i {
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
		file_apps_code_pb_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCodeResponse); i {
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
		file_apps_code_pb_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRequest); i {
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
			RawDescriptor: file_apps_code_pb_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_code_pb_code_proto_goTypes,
		DependencyIndexes: file_apps_code_pb_code_proto_depIdxs,
		EnumInfos:         file_apps_code_pb_code_proto_enumTypes,
		MessageInfos:      file_apps_code_pb_code_proto_msgTypes,
	}.Build()
	File_apps_code_pb_code_proto = out.File
	file_apps_code_pb_code_proto_rawDesc = nil
	file_apps_code_pb_code_proto_goTypes = nil
	file_apps_code_pb_code_proto_depIdxs = nil
}
