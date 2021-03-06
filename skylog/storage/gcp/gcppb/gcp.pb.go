// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: gcp.proto

package gcppb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// BuildJob is a Merke tree building job. It instructs workers to build a
// subtree covering leaves of the [begin, end) range for the specified tree.
type BuildJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tree ID to build a subtree of.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	// The beginning of the leaves range (inclusive).
	Begin uint64 `protobuf:"varint,2,opt,name=begin,proto3" json:"begin,omitempty"`
	// The ending of the leaves range (exclusive).
	End uint64 `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	// The tree sharding scheme that the build worker should use. Since builders
	// are stateless it is more efficient to pass it in from the master rather
	// than fetch it from storage on each job invocation.
	TreeSharding *TreeSharding `protobuf:"bytes,4,opt,name=tree_sharding,json=treeSharding,proto3" json:"tree_sharding,omitempty"`
	// The sequence sharding scheme that the build worker should use. It is
	// provided for the same reason as the tree sharding field.
	SeqSharding *SequenceSharding `protobuf:"bytes,5,opt,name=seq_sharding,json=seqSharding,proto3" json:"seq_sharding,omitempty"`
}

func (x *BuildJob) Reset() {
	*x = BuildJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gcp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildJob) ProtoMessage() {}

func (x *BuildJob) ProtoReflect() protoreflect.Message {
	mi := &file_gcp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildJob.ProtoReflect.Descriptor instead.
func (*BuildJob) Descriptor() ([]byte, []int) {
	return file_gcp_proto_rawDescGZIP(), []int{0}
}

func (x *BuildJob) GetTreeId() int64 {
	if x != nil {
		return x.TreeId
	}
	return 0
}

func (x *BuildJob) GetBegin() uint64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *BuildJob) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *BuildJob) GetTreeSharding() *TreeSharding {
	if x != nil {
		return x.TreeSharding
	}
	return nil
}

func (x *BuildJob) GetSeqSharding() *SequenceSharding {
	if x != nil {
		return x.SeqSharding
	}
	return nil
}

// TreeSharding describes the tree nodes sharding scheme.
//
// The specified number of lower tree levels are split into the specified number
// of shards, where each shard stores a periodic sub-structure of perfect
// subtrees. There is one extra shard covering the upper tree levels.
//
// See the tree storage comments for more details.
type TreeSharding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Levels uint32 `protobuf:"varint,1,opt,name=levels,proto3" json:"levels,omitempty"`
	Shards uint32 `protobuf:"varint,2,opt,name=shards,proto3" json:"shards,omitempty"`
}

func (x *TreeSharding) Reset() {
	*x = TreeSharding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gcp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreeSharding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreeSharding) ProtoMessage() {}

func (x *TreeSharding) ProtoReflect() protoreflect.Message {
	mi := &file_gcp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreeSharding.ProtoReflect.Descriptor instead.
func (*TreeSharding) Descriptor() ([]byte, []int) {
	return file_gcp_proto_rawDescGZIP(), []int{1}
}

func (x *TreeSharding) GetLevels() uint32 {
	if x != nil {
		return x.Levels
	}
	return 0
}

func (x *TreeSharding) GetShards() uint32 {
	if x != nil {
		return x.Shards
	}
	return 0
}

// SequenceSharding describes the sequenced log entries sharding scheme.
//
// The sequence is split into the specified number of shards, where each shard
// stores a periodic sub-sequence consisting of stripes of the specified size.
//
// See the sequence storage comments for more details.
type SequenceSharding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shards uint32 `protobuf:"varint,1,opt,name=shards,proto3" json:"shards,omitempty"`
	Size   uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SequenceSharding) Reset() {
	*x = SequenceSharding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gcp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceSharding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceSharding) ProtoMessage() {}

func (x *SequenceSharding) ProtoReflect() protoreflect.Message {
	mi := &file_gcp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceSharding.ProtoReflect.Descriptor instead.
func (*SequenceSharding) Descriptor() ([]byte, []int) {
	return file_gcp_proto_rawDescGZIP(), []int{2}
}

func (x *SequenceSharding) GetShards() uint32 {
	if x != nil {
		return x.Shards
	}
	return 0
}

func (x *SequenceSharding) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_gcp_proto protoreflect.FileDescriptor

var file_gcp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x63, 0x70,
	0x70, 0x62, 0x22, 0xc1, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4a, 0x6f, 0x62, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x72, 0x65, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x38, 0x0a, 0x0d, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x63, 0x70, 0x70, 0x62, 0x2e,
	0x54, 0x72, 0x65, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x74, 0x72,
	0x65, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x65,
	0x71, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x63, 0x70, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x73, 0x65, 0x71, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x0c, 0x54, 0x72, 0x65, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x6c,
	0x6c, 0x69, 0x61, 0x6e, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gcp_proto_rawDescOnce sync.Once
	file_gcp_proto_rawDescData = file_gcp_proto_rawDesc
)

func file_gcp_proto_rawDescGZIP() []byte {
	file_gcp_proto_rawDescOnce.Do(func() {
		file_gcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_gcp_proto_rawDescData)
	})
	return file_gcp_proto_rawDescData
}

var file_gcp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gcp_proto_goTypes = []interface{}{
	(*BuildJob)(nil),         // 0: gcppb.BuildJob
	(*TreeSharding)(nil),     // 1: gcppb.TreeSharding
	(*SequenceSharding)(nil), // 2: gcppb.SequenceSharding
}
var file_gcp_proto_depIdxs = []int32{
	1, // 0: gcppb.BuildJob.tree_sharding:type_name -> gcppb.TreeSharding
	2, // 1: gcppb.BuildJob.seq_sharding:type_name -> gcppb.SequenceSharding
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gcp_proto_init() }
func file_gcp_proto_init() {
	if File_gcp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gcp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildJob); i {
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
		file_gcp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreeSharding); i {
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
		file_gcp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceSharding); i {
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
			RawDescriptor: file_gcp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gcp_proto_goTypes,
		DependencyIndexes: file_gcp_proto_depIdxs,
		MessageInfos:      file_gcp_proto_msgTypes,
	}.Build()
	File_gcp_proto = out.File
	file_gcp_proto_rawDesc = nil
	file_gcp_proto_goTypes = nil
	file_gcp_proto_depIdxs = nil
}
