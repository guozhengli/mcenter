// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/code/pb/rpc.proto

package code

import (
	token "github.com/infraboard/mcenter/apps/token"
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

// IssueCodeRequest 验证码申请请求
type IssueCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 颁发方式
	// @gotags: json:"issue_by"
	IssueBy token.GRANT_TYPE `protobuf:"varint,1,opt,name=issue_by,json=issueBy,proto3,enum=infraboard.mcenter.token.GRANT_TYPE" json:"issue_by"`
	// PASSWORD/LDAP授权时, 用户名
	// @gotags: json:"username,omitempty"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// PASSWORD/LDAP授权时, 用户密码
	// @gotags: json:"password,omitempty"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// PRIVATE_TOKEN授权时, 访问令牌
	// @gotags: json:"access_token,omitempty"
	AccessToken string `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// AUTH_CODE授权时, Code, 针对第三方登陆: 飞书/钉钉/企业微信
	// @gotags: json:"auth_code"
	AuthCode string `protobuf:"bytes,5,opt,name=auth_code,json=authCode,proto3" json:"auth_code"`
}

func (x *IssueCodeRequest) Reset() {
	*x = IssueCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_code_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeRequest) ProtoMessage() {}

func (x *IssueCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_rpc_proto_msgTypes[0]
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
	return file_apps_code_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IssueCodeRequest) GetIssueBy() token.GRANT_TYPE {
	if x != nil {
		return x.IssueBy
	}
	return token.GRANT_TYPE(0)
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

func (x *IssueCodeRequest) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
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
		mi := &file_apps_code_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeResponse) ProtoMessage() {}

func (x *IssueCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_rpc_proto_msgTypes[1]
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
	return file_apps_code_pb_rpc_proto_rawDescGZIP(), []int{1}
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
		mi := &file_apps_code_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRequest) ProtoMessage() {}

func (x *VerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_code_pb_rpc_proto_msgTypes[2]
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
	return file_apps_code_pb_rpc_proto_rawDescGZIP(), []int{2}
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

var File_apps_code_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_code_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xc2, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x62, 0x0a, 0x09, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_code_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_code_pb_rpc_proto_rawDescData = file_apps_code_pb_rpc_proto_rawDesc
)

func file_apps_code_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_code_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_code_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_code_pb_rpc_proto_rawDescData)
	})
	return file_apps_code_pb_rpc_proto_rawDescData
}

var file_apps_code_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_code_pb_rpc_proto_goTypes = []interface{}{
	(*IssueCodeRequest)(nil),  // 0: infraboard.mcenter.code.IssueCodeRequest
	(*IssueCodeResponse)(nil), // 1: infraboard.mcenter.code.IssueCodeResponse
	(*VerifyCodeRequest)(nil), // 2: infraboard.mcenter.code.VerifyCodeRequest
	(token.GRANT_TYPE)(0),     // 3: infraboard.mcenter.token.GRANT_TYPE
	(*Code)(nil),              // 4: infraboard.mcenter.code.Code
}
var file_apps_code_pb_rpc_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.code.IssueCodeRequest.issue_by:type_name -> infraboard.mcenter.token.GRANT_TYPE
	0, // 1: infraboard.mcenter.code.RPC.IssueCode:input_type -> infraboard.mcenter.code.IssueCodeRequest
	2, // 2: infraboard.mcenter.code.RPC.VerifyCode:input_type -> infraboard.mcenter.code.VerifyCodeRequest
	1, // 3: infraboard.mcenter.code.RPC.IssueCode:output_type -> infraboard.mcenter.code.IssueCodeResponse
	4, // 4: infraboard.mcenter.code.RPC.VerifyCode:output_type -> infraboard.mcenter.code.Code
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_code_pb_rpc_proto_init() }
func file_apps_code_pb_rpc_proto_init() {
	if File_apps_code_pb_rpc_proto != nil {
		return
	}
	file_apps_code_pb_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_code_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_code_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_code_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_apps_code_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_code_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_code_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_code_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_code_pb_rpc_proto = out.File
	file_apps_code_pb_rpc_proto_rawDesc = nil
	file_apps_code_pb_rpc_proto_goTypes = nil
	file_apps_code_pb_rpc_proto_depIdxs = nil
}
