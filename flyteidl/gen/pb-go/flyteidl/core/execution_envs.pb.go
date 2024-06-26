// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/core/execution_envs.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ExecutionEnvAssignment is a message that is used to assign an execution environment to a set of
// nodes.
type ExecutionEnvAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// node_ids is a list of node ids that are being assigned the execution environment.
	NodeIds []string `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	// task_type is the type of task that is being assigned. This is used to override which Flyte
	// plugin will be used during execution.
	TaskType string `protobuf:"bytes,2,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// execution_env is the environment that is being assigned to the nodes.
	ExecutionEnv *ExecutionEnv `protobuf:"bytes,3,opt,name=execution_env,json=executionEnv,proto3" json:"execution_env,omitempty"`
}

func (x *ExecutionEnvAssignment) Reset() {
	*x = ExecutionEnvAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_execution_envs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionEnvAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionEnvAssignment) ProtoMessage() {}

func (x *ExecutionEnvAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_execution_envs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionEnvAssignment.ProtoReflect.Descriptor instead.
func (*ExecutionEnvAssignment) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_execution_envs_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionEnvAssignment) GetNodeIds() []string {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

func (x *ExecutionEnvAssignment) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *ExecutionEnvAssignment) GetExecutionEnv() *ExecutionEnv {
	if x != nil {
		return x.ExecutionEnv
	}
	return nil
}

// ExecutionEnv is a message that is used to specify the execution environment.
type ExecutionEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is a human-readable identifier for the execution environment. This is combined with the
	// project, domain, and version to uniquely identify an execution environment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type is the type of the execution environment.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// environment is a oneof field that can be used to specify the environment in different ways.
	//
	// Types that are assignable to Environment:
	//
	//	*ExecutionEnv_Extant
	//	*ExecutionEnv_Spec
	Environment isExecutionEnv_Environment `protobuf_oneof:"environment"`
	// version is the version of the execution environment. This may be used differently by each
	// individual environment type (ex. auto-generated or manually provided), but is intended to
	// allow variance in environment specifications with the same ID.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ExecutionEnv) Reset() {
	*x = ExecutionEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_execution_envs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionEnv) ProtoMessage() {}

func (x *ExecutionEnv) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_execution_envs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionEnv.ProtoReflect.Descriptor instead.
func (*ExecutionEnv) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_execution_envs_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionEnv) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExecutionEnv) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (m *ExecutionEnv) GetEnvironment() isExecutionEnv_Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (x *ExecutionEnv) GetExtant() *structpb.Struct {
	if x, ok := x.GetEnvironment().(*ExecutionEnv_Extant); ok {
		return x.Extant
	}
	return nil
}

func (x *ExecutionEnv) GetSpec() *structpb.Struct {
	if x, ok := x.GetEnvironment().(*ExecutionEnv_Spec); ok {
		return x.Spec
	}
	return nil
}

func (x *ExecutionEnv) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type isExecutionEnv_Environment interface {
	isExecutionEnv_Environment()
}

type ExecutionEnv_Extant struct {
	// extant is a reference to an existing environment.
	Extant *structpb.Struct `protobuf:"bytes,3,opt,name=extant,proto3,oneof"`
}

type ExecutionEnv_Spec struct {
	// spec is a specification of the environment.
	Spec *structpb.Struct `protobuf:"bytes,4,opt,name=spec,proto3,oneof"`
}

func (*ExecutionEnv_Extant) isExecutionEnv_Environment() {}

func (*ExecutionEnv_Spec) isExecutionEnv_Environment() {}

var File_flyteidl_core_execution_envs_proto protoreflect.FileDescriptor

var file_flyteidl_core_execution_envs_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x76, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x76, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x31, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xb8, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x42, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a,
	0x3a, 0x43, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_execution_envs_proto_rawDescOnce sync.Once
	file_flyteidl_core_execution_envs_proto_rawDescData = file_flyteidl_core_execution_envs_proto_rawDesc
)

func file_flyteidl_core_execution_envs_proto_rawDescGZIP() []byte {
	file_flyteidl_core_execution_envs_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_execution_envs_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_execution_envs_proto_rawDescData)
	})
	return file_flyteidl_core_execution_envs_proto_rawDescData
}

var file_flyteidl_core_execution_envs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flyteidl_core_execution_envs_proto_goTypes = []interface{}{
	(*ExecutionEnvAssignment)(nil), // 0: flyteidl.core.ExecutionEnvAssignment
	(*ExecutionEnv)(nil),           // 1: flyteidl.core.ExecutionEnv
	(*structpb.Struct)(nil),        // 2: google.protobuf.Struct
}
var file_flyteidl_core_execution_envs_proto_depIdxs = []int32{
	1, // 0: flyteidl.core.ExecutionEnvAssignment.execution_env:type_name -> flyteidl.core.ExecutionEnv
	2, // 1: flyteidl.core.ExecutionEnv.extant:type_name -> google.protobuf.Struct
	2, // 2: flyteidl.core.ExecutionEnv.spec:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flyteidl_core_execution_envs_proto_init() }
func file_flyteidl_core_execution_envs_proto_init() {
	if File_flyteidl_core_execution_envs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_execution_envs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionEnvAssignment); i {
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
		file_flyteidl_core_execution_envs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionEnv); i {
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
	file_flyteidl_core_execution_envs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExecutionEnv_Extant)(nil),
		(*ExecutionEnv_Spec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_core_execution_envs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_execution_envs_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_execution_envs_proto_depIdxs,
		MessageInfos:      file_flyteidl_core_execution_envs_proto_msgTypes,
	}.Build()
	File_flyteidl_core_execution_envs_proto = out.File
	file_flyteidl_core_execution_envs_proto_rawDesc = nil
	file_flyteidl_core_execution_envs_proto_goTypes = nil
	file_flyteidl_core_execution_envs_proto_depIdxs = nil
}
