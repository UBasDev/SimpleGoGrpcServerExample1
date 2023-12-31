// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: author1.proto

package goproject3

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

type AuthorCreateRequest1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorName string `protobuf:"bytes,1,opt,name=authorName,proto3" json:"authorName,omitempty"`
	AuthorAge  int64  `protobuf:"varint,2,opt,name=authorAge,proto3" json:"authorAge,omitempty"`
}

func (x *AuthorCreateRequest1) Reset() {
	*x = AuthorCreateRequest1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorCreateRequest1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorCreateRequest1) ProtoMessage() {}

func (x *AuthorCreateRequest1) ProtoReflect() protoreflect.Message {
	mi := &file_author1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorCreateRequest1.ProtoReflect.Descriptor instead.
func (*AuthorCreateRequest1) Descriptor() ([]byte, []int) {
	return file_author1_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorCreateRequest1) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *AuthorCreateRequest1) GetAuthorAge() int64 {
	if x != nil {
		return x.AuthorAge
	}
	return 0
}

type AuthorCreateResponse1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId   int64  `protobuf:"varint,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	AuthorName string `protobuf:"bytes,2,opt,name=authorName,proto3" json:"authorName,omitempty"`
	AuthorAge  int64  `protobuf:"varint,3,opt,name=authorAge,proto3" json:"authorAge,omitempty"`
}

func (x *AuthorCreateResponse1) Reset() {
	*x = AuthorCreateResponse1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorCreateResponse1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorCreateResponse1) ProtoMessage() {}

func (x *AuthorCreateResponse1) ProtoReflect() protoreflect.Message {
	mi := &file_author1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorCreateResponse1.ProtoReflect.Descriptor instead.
func (*AuthorCreateResponse1) Descriptor() ([]byte, []int) {
	return file_author1_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorCreateResponse1) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *AuthorCreateResponse1) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *AuthorCreateResponse1) GetAuthorAge() int64 {
	if x != nil {
		return x.AuthorAge
	}
	return 0
}

var File_author1_proto protoreflect.FileDescriptor

var file_author1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x54, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x41, 0x67, 0x65, 0x22, 0x71, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x41, 0x67, 0x65, 0x32, 0x50, 0x0a, 0x07, 0x41, 0x75, 0x68, 0x74,
	0x6f, 0x72, 0x31, 0x12, 0x45, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x31, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x31, 0x1a, 0x16, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x42, 0x18, 0x5a, 0x16, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_author1_proto_rawDescOnce sync.Once
	file_author1_proto_rawDescData = file_author1_proto_rawDesc
)

func file_author1_proto_rawDescGZIP() []byte {
	file_author1_proto_rawDescOnce.Do(func() {
		file_author1_proto_rawDescData = protoimpl.X.CompressGZIP(file_author1_proto_rawDescData)
	})
	return file_author1_proto_rawDescData
}

var file_author1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_author1_proto_goTypes = []interface{}{
	(*AuthorCreateRequest1)(nil),  // 0: AuthorCreateRequest1
	(*AuthorCreateResponse1)(nil), // 1: AuthorCreateResponse1
}
var file_author1_proto_depIdxs = []int32{
	0, // 0: Auhtor1.Author1CreateService:input_type -> AuthorCreateRequest1
	1, // 1: Auhtor1.Author1CreateService:output_type -> AuthorCreateResponse1
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_author1_proto_init() }
func file_author1_proto_init() {
	if File_author1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_author1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorCreateRequest1); i {
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
		file_author1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorCreateResponse1); i {
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
			RawDescriptor: file_author1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_author1_proto_goTypes,
		DependencyIndexes: file_author1_proto_depIdxs,
		MessageInfos:      file_author1_proto_msgTypes,
	}.Build()
	File_author1_proto = out.File
	file_author1_proto_rawDesc = nil
	file_author1_proto_goTypes = nil
	file_author1_proto_depIdxs = nil
}
