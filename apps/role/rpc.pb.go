// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/role/pb/rpc.proto

package role

import (
	request "github.com/infraboard/mcube/http/request"
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

// QueryPermissionRequest 查询用户权限
type QueryPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace"`
	// @gotags: json:"username"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
	// 查询角色权限
	// @gotags: json:"role_id"
	RoleId string `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	// 忽略数据
	// @gotags: json:"skip_items"
	SkipItems bool `protobuf:"varint,5,opt,name=skip_items,json=skipItems,proto3" json:"skip_items"`
}

func (x *QueryPermissionRequest) Reset() {
	*x = QueryPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPermissionRequest) ProtoMessage() {}

func (x *QueryPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPermissionRequest.ProtoReflect.Descriptor instead.
func (*QueryPermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPermissionRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPermissionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *QueryPermissionRequest) GetSkipItems() bool {
	if x != nil {
		return x.SkipItems
	}
	return false
}

// QueryRoleRequest 列表查询
type QueryRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"type"
	Type *RoleType `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.mcenter.role.RoleType,oneof" json:"type"`
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain"`
	// @gotags: json:"with_permission"
	WithPermission bool `protobuf:"varint,4,opt,name=with_permission,json=withPermission,proto3" json:"with_permission"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRoleRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRoleRequest) GetType() RoleType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return RoleType_BUILDIN
}

func (x *QueryRoleRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryRoleRequest) GetWithPermission() bool {
	if x != nil {
		return x.WithPermission
	}
	return false
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @gotags: json:"name,omitempty" validate:"required,lte=64"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required,lte=64"`
	// @gotags: bson:"type" json:"type"
	Type RoleType `protobuf:"varint,3,opt,name=type,proto3,enum=infraboard.mcenter.role.RoleType" json:"type" bson:"type"`
}

func (x *DescribeRoleRequest) Reset() {
	*x = DescribeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRoleRequest) ProtoMessage() {}

func (x *DescribeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRoleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_BUILDIN
}

// DescribeRoleRequest role详情
type DescribePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePermissionRequest) Reset() {
	*x = DescribePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePermissionRequest) ProtoMessage() {}

func (x *DescribePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePermissionRequest.ProtoReflect.Descriptor instead.
func (*DescribePermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// CheckPermissionRequest todo
type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace"`
	// @gotags: json:"endpoint_id"
	EndpointId string `protobuf:"bytes,3,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id"`
	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
	// @gotags: json:"path"
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`
	// @gotags: json:"username"
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *CheckPermissionRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *CheckPermissionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CheckPermissionRequest) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *CheckPermissionRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CheckPermissionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckPermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_apps_role_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_role_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x16, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd0,
	0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x70, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xde, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0x97, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x58, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x5b, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x6a, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x12,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_role_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_role_pb_rpc_proto_rawDescData = file_apps_role_pb_rpc_proto_rawDesc
)

func file_apps_role_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_role_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_role_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_role_pb_rpc_proto_rawDescData)
	})
	return file_apps_role_pb_rpc_proto_rawDescData
}

var file_apps_role_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_role_pb_rpc_proto_goTypes = []interface{}{
	(*QueryPermissionRequest)(nil),    // 0: infraboard.mcenter.role.QueryPermissionRequest
	(*QueryRoleRequest)(nil),          // 1: infraboard.mcenter.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil),       // 2: infraboard.mcenter.role.DescribeRoleRequest
	(*DescribePermissionRequest)(nil), // 3: infraboard.mcenter.role.DescribePermissionRequest
	(*CheckPermissionRequest)(nil),    // 4: infraboard.mcenter.role.CheckPermissionRequest
	(*request.PageRequest)(nil),       // 5: infraboard.mcube.page.PageRequest
	(RoleType)(0),                     // 6: infraboard.mcenter.role.RoleType
	(*RoleSet)(nil),                   // 7: infraboard.mcenter.role.RoleSet
	(*Role)(nil),                      // 8: infraboard.mcenter.role.Role
	(*PermissionSet)(nil),             // 9: infraboard.mcenter.role.PermissionSet
	(*Permission)(nil),                // 10: infraboard.mcenter.role.Permission
}
var file_apps_role_pb_rpc_proto_depIdxs = []int32{
	5,  // 0: infraboard.mcenter.role.QueryPermissionRequest.page:type_name -> infraboard.mcube.page.PageRequest
	5,  // 1: infraboard.mcenter.role.QueryRoleRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6,  // 2: infraboard.mcenter.role.QueryRoleRequest.type:type_name -> infraboard.mcenter.role.RoleType
	6,  // 3: infraboard.mcenter.role.DescribeRoleRequest.type:type_name -> infraboard.mcenter.role.RoleType
	5,  // 4: infraboard.mcenter.role.CheckPermissionRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1,  // 5: infraboard.mcenter.role.RPC.QueryRole:input_type -> infraboard.mcenter.role.QueryRoleRequest
	2,  // 6: infraboard.mcenter.role.RPC.DescribeRole:input_type -> infraboard.mcenter.role.DescribeRoleRequest
	0,  // 7: infraboard.mcenter.role.RPC.QueryPermission:input_type -> infraboard.mcenter.role.QueryPermissionRequest
	3,  // 8: infraboard.mcenter.role.RPC.DescribePermission:input_type -> infraboard.mcenter.role.DescribePermissionRequest
	7,  // 9: infraboard.mcenter.role.RPC.QueryRole:output_type -> infraboard.mcenter.role.RoleSet
	8,  // 10: infraboard.mcenter.role.RPC.DescribeRole:output_type -> infraboard.mcenter.role.Role
	9,  // 11: infraboard.mcenter.role.RPC.QueryPermission:output_type -> infraboard.mcenter.role.PermissionSet
	10, // 12: infraboard.mcenter.role.RPC.DescribePermission:output_type -> infraboard.mcenter.role.Permission
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_role_pb_rpc_proto_init() }
func file_apps_role_pb_rpc_proto_init() {
	if File_apps_role_pb_rpc_proto != nil {
		return
	}
	file_apps_role_pb_role_proto_init()
	file_apps_role_pb_permission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_role_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPermissionRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRoleRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePermissionRequest); i {
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
		file_apps_role_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
	file_apps_role_pb_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_role_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_role_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_role_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_role_pb_rpc_proto = out.File
	file_apps_role_pb_rpc_proto_rawDesc = nil
	file_apps_role_pb_rpc_proto_goTypes = nil
	file_apps_role_pb_rpc_proto_depIdxs = nil
}
