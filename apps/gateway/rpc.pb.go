// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/gateway/pb/rpc.proto

package gateway

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

// QueryRoleRequest 列表查询
type QueryGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain"`
}

func (x *QueryGatewayRequest) Reset() {
	*x = QueryGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGatewayRequest) ProtoMessage() {}

func (x *QueryGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGatewayRequest.ProtoReflect.Descriptor instead.
func (*QueryGatewayRequest) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryGatewayRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryGatewayRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DescribeRoleRequest role详情
type DescribeGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeGatewayRequest) Reset() {
	*x = DescribeGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeGatewayRequest) ProtoMessage() {}

func (x *DescribeGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeGatewayRequest.ProtoReflect.Descriptor instead.
func (*DescribeGatewayRequest) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeGatewayRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_gateway_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_gateway_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x28, 0x0a,
	0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xda, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x67, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x74, 0x12, 0x6a, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x32, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_gateway_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_gateway_pb_rpc_proto_rawDescData = file_apps_gateway_pb_rpc_proto_rawDesc
)

func file_apps_gateway_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_gateway_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_gateway_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_gateway_pb_rpc_proto_rawDescData)
	})
	return file_apps_gateway_pb_rpc_proto_rawDescData
}

var file_apps_gateway_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_gateway_pb_rpc_proto_goTypes = []interface{}{
	(*QueryGatewayRequest)(nil),    // 0: infraboard.mcenter.gateway.QueryGatewayRequest
	(*DescribeGatewayRequest)(nil), // 1: infraboard.mcenter.gateway.DescribeGatewayRequest
	(*request.PageRequest)(nil),    // 2: infraboard.mcube.page.PageRequest
	(*GatewaySet)(nil),             // 3: infraboard.mcenter.gateway.GatewaySet
	(*Gateway)(nil),                // 4: infraboard.mcenter.gateway.Gateway
}
var file_apps_gateway_pb_rpc_proto_depIdxs = []int32{
	2, // 0: infraboard.mcenter.gateway.QueryGatewayRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 1: infraboard.mcenter.gateway.RPC.QueryGateway:input_type -> infraboard.mcenter.gateway.QueryGatewayRequest
	1, // 2: infraboard.mcenter.gateway.RPC.DescribeGateway:input_type -> infraboard.mcenter.gateway.DescribeGatewayRequest
	3, // 3: infraboard.mcenter.gateway.RPC.QueryGateway:output_type -> infraboard.mcenter.gateway.GatewaySet
	4, // 4: infraboard.mcenter.gateway.RPC.DescribeGateway:output_type -> infraboard.mcenter.gateway.Gateway
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_gateway_pb_rpc_proto_init() }
func file_apps_gateway_pb_rpc_proto_init() {
	if File_apps_gateway_pb_rpc_proto != nil {
		return
	}
	file_apps_gateway_pb_gateway_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_gateway_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGatewayRequest); i {
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
		file_apps_gateway_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeGatewayRequest); i {
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
			RawDescriptor: file_apps_gateway_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_gateway_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_gateway_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_gateway_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_gateway_pb_rpc_proto = out.File
	file_apps_gateway_pb_rpc_proto_rawDesc = nil
	file_apps_gateway_pb_rpc_proto_goTypes = nil
	file_apps_gateway_pb_rpc_proto_depIdxs = nil
}
