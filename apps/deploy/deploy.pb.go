// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: apps/deploy/pb/deploy.proto

package deploy

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

type TYPE int32

const (
	// 基于k8s部署
	TYPE_KUBERNETES TYPE = 0
	// 基于主机部署
	TYPE_HOST TYPE = 1
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "KUBERNETES",
		1: "HOST",
	}
	TYPE_value = map[string]int32{
		"KUBERNETES": 0,
		"HOST":       1,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_deploy_pb_deploy_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_deploy_pb_deploy_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_deploy_pb_deploy_proto_rawDescGZIP(), []int{0}
}

type Deploy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署Id
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 创建信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateDeployRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Deploy) Reset() {
	*x = Deploy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_deploy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deploy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deploy) ProtoMessage() {}

func (x *Deploy) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_deploy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deploy.ProtoReflect.Descriptor instead.
func (*Deploy) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_deploy_proto_rawDescGZIP(), []int{0}
}

func (x *Deploy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deploy) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Deploy) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Deploy) GetSpec() *CreateDeployRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

// K8sTypeConfig yaml文本格式的k8s部署相关配置文件
type K8STypeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 部署集群的名称
	// @gotags: bson:"cluster_name" json:"cluster_name" validate:"required"
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name" bson:"cluster_name" validate:"required"`
	// k8s yaml配置, 支持deploy/statfulset/daemonset/job/cronjob
	// @gotags: bson:"workload" json:"workload" validate:"required"
	Workload string `protobuf:"bytes,2,opt,name=workload,proto3" json:"workload" bson:"workload" validate:"required"`
	// k8s service配置
	// @gotags: bson:"service" json:"service" validate:"required"
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service" bson:"service" validate:"required"`
}

func (x *K8STypeConfig) Reset() {
	*x = K8STypeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_deploy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8STypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8STypeConfig) ProtoMessage() {}

func (x *K8STypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_deploy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8STypeConfig.ProtoReflect.Descriptor instead.
func (*K8STypeConfig) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_deploy_proto_rawDescGZIP(), []int{1}
}

func (x *K8STypeConfig) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *K8STypeConfig) GetWorkload() string {
	if x != nil {
		return x.Workload
	}
	return ""
}

func (x *K8STypeConfig) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type CreateDeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 那个服务的部署
	// @gotags: bson:"service_id" json:"service_id"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// 部署环境
	// @gotags: bson:"environment" json:"environment" validate:"required"
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment" bson:"environment" validate:"required"`
	// 部署资源提供方
	// @gotags: bson:"provider" json:"provider" validate:"required"
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider" bson:"provider" validate:"required"`
	// 部署地域
	// @gotags: bson:"region" json:"region" validate:"required"
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region" bson:"region" validate:"required"`
	// 部署方式
	// @gotags: bson:"type" json:"type"
	Type TYPE `protobuf:"varint,5,opt,name=type,proto3,enum=infraboard.mcenter.deploy.TYPE" json:"type" bson:"type"`
	// k8s模式下的部署配置
	// @gotags: json:"k8s_type_config" bson:"k8s_type_config"
	K8STypeConfig *K8STypeConfig `protobuf:"bytes,11,opt,name=k8s_type_config,json=k8sTypeConfig,proto3" json:"k8s_type_config" bson:"k8s_type_config"`
	// 部署的名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" bson:"name"`
	// 部署描述信息
	// @gotags: bson:"describe" json:"describe"
	Describe string `protobuf:"bytes,7,opt,name=describe,proto3" json:"describe" bson:"describe"`
	// 部署标签
	// @gotags: bson:"tags" json:"tags"
	Tags map[string]string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"tags"`
}

func (x *CreateDeployRequest) Reset() {
	*x = CreateDeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_deploy_pb_deploy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeployRequest) ProtoMessage() {}

func (x *CreateDeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_deploy_pb_deploy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeployRequest.ProtoReflect.Descriptor instead.
func (*CreateDeployRequest) Descriptor() ([]byte, []int) {
	return file_apps_deploy_pb_deploy_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeployRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreateDeployRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *CreateDeployRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateDeployRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateDeployRequest) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_KUBERNETES
}

func (x *CreateDeployRequest) GetK8STypeConfig() *K8STypeConfig {
	if x != nil {
		return x.K8STypeConfig
	}
	return nil
}

func (x *CreateDeployRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDeployRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *CreateDeployRequest) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_apps_deploy_pb_deploy_proto protoreflect.FileDescriptor

var file_apps_deploy_pb_deploy_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x42, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x22, 0x68, 0x0a, 0x0d, 0x4b, 0x38, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xc8, 0x03, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a,
	0x0f, 0x6b, 0x38, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x2e, 0x4b, 0x38, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0d, 0x6b, 0x38, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x4c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a,
	0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x20, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0e,
	0x0a, 0x0a, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_deploy_pb_deploy_proto_rawDescOnce sync.Once
	file_apps_deploy_pb_deploy_proto_rawDescData = file_apps_deploy_pb_deploy_proto_rawDesc
)

func file_apps_deploy_pb_deploy_proto_rawDescGZIP() []byte {
	file_apps_deploy_pb_deploy_proto_rawDescOnce.Do(func() {
		file_apps_deploy_pb_deploy_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_deploy_pb_deploy_proto_rawDescData)
	})
	return file_apps_deploy_pb_deploy_proto_rawDescData
}

var file_apps_deploy_pb_deploy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_deploy_pb_deploy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_deploy_pb_deploy_proto_goTypes = []interface{}{
	(TYPE)(0),                   // 0: infraboard.mcenter.deploy.TYPE
	(*Deploy)(nil),              // 1: infraboard.mcenter.deploy.Deploy
	(*K8STypeConfig)(nil),       // 2: infraboard.mcenter.deploy.K8sTypeConfig
	(*CreateDeployRequest)(nil), // 3: infraboard.mcenter.deploy.CreateDeployRequest
	nil,                         // 4: infraboard.mcenter.deploy.CreateDeployRequest.TagsEntry
}
var file_apps_deploy_pb_deploy_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.deploy.Deploy.spec:type_name -> infraboard.mcenter.deploy.CreateDeployRequest
	0, // 1: infraboard.mcenter.deploy.CreateDeployRequest.type:type_name -> infraboard.mcenter.deploy.TYPE
	2, // 2: infraboard.mcenter.deploy.CreateDeployRequest.k8s_type_config:type_name -> infraboard.mcenter.deploy.K8sTypeConfig
	4, // 3: infraboard.mcenter.deploy.CreateDeployRequest.tags:type_name -> infraboard.mcenter.deploy.CreateDeployRequest.TagsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_deploy_pb_deploy_proto_init() }
func file_apps_deploy_pb_deploy_proto_init() {
	if File_apps_deploy_pb_deploy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_deploy_pb_deploy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deploy); i {
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
		file_apps_deploy_pb_deploy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8STypeConfig); i {
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
		file_apps_deploy_pb_deploy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeployRequest); i {
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
			RawDescriptor: file_apps_deploy_pb_deploy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_deploy_pb_deploy_proto_goTypes,
		DependencyIndexes: file_apps_deploy_pb_deploy_proto_depIdxs,
		EnumInfos:         file_apps_deploy_pb_deploy_proto_enumTypes,
		MessageInfos:      file_apps_deploy_pb_deploy_proto_msgTypes,
	}.Build()
	File_apps_deploy_pb_deploy_proto = out.File
	file_apps_deploy_pb_deploy_proto_rawDesc = nil
	file_apps_deploy_pb_deploy_proto_goTypes = nil
	file_apps_deploy_pb_deploy_proto_depIdxs = nil
}
