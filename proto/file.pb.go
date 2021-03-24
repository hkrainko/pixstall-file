//cmd: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/file.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/file.proto

package proto

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

type SaveFileRequest_FileType int32

const (
	SaveFileRequest_Message          SaveFileRequest_FileType = 0
	SaveFileRequest_Completion       SaveFileRequest_FileType = 1
	SaveFileRequest_CommissionRefImg SaveFileRequest_FileType = 2
	SaveFileRequest_Artwork          SaveFileRequest_FileType = 3
	SaveFileRequest_Roof             SaveFileRequest_FileType = 4
	SaveFileRequest_OpenCommission   SaveFileRequest_FileType = 5
	SaveFileRequest_Profile          SaveFileRequest_FileType = 6
)

// Enum value maps for SaveFileRequest_FileType.
var (
	SaveFileRequest_FileType_name = map[int32]string{
		0: "Message",
		1: "Completion",
		2: "CommissionRefImg",
		3: "Artwork",
		4: "Roof",
		5: "OpenCommission",
		6: "Profile",
	}
	SaveFileRequest_FileType_value = map[string]int32{
		"Message":          0,
		"Completion":       1,
		"CommissionRefImg": 2,
		"Artwork":          3,
		"Roof":             4,
		"OpenCommission":   5,
		"Profile":          6,
	}
)

func (x SaveFileRequest_FileType) Enum() *SaveFileRequest_FileType {
	p := new(SaveFileRequest_FileType)
	*p = x
	return p
}

func (x SaveFileRequest_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SaveFileRequest_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_file_proto_enumTypes[0].Descriptor()
}

func (SaveFileRequest_FileType) Type() protoreflect.EnumType {
	return &file_proto_file_proto_enumTypes[0]
}

func (x SaveFileRequest_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SaveFileRequest_FileType.Descriptor instead.
func (SaveFileRequest_FileType) EnumDescriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{0, 0}
}

type SaveFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     []byte                   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	FileType SaveFileRequest_FileType `protobuf:"varint,2,opt,name=fileType,proto3,enum=SaveFileRequest_FileType" json:"fileType,omitempty"`
}

func (x *SaveFileRequest) Reset() {
	*x = SaveFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileRequest) ProtoMessage() {}

func (x *SaveFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileRequest.ProtoReflect.Descriptor instead.
func (*SaveFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{0}
}

func (x *SaveFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *SaveFileRequest) GetFileType() SaveFileRequest_FileType {
	if x != nil {
		return x.FileType
	}
	return SaveFileRequest_Message
}

type SaveFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *SaveFileResponse) Reset() {
	*x = SaveFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileResponse) ProtoMessage() {}

func (x *SaveFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileResponse.ProtoReflect.Descriptor instead.
func (*SaveFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{1}
}

func (x *SaveFileResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_proto_file_proto protoreflect.FileDescriptor

var file_proto_file_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x75, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x49, 0x6d, 0x67, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x6f, 0x6f, 0x66, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x10, 0x06, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x32, 0x40, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x31, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_file_proto_rawDescOnce sync.Once
	file_proto_file_proto_rawDescData = file_proto_file_proto_rawDesc
)

func file_proto_file_proto_rawDescGZIP() []byte {
	file_proto_file_proto_rawDescOnce.Do(func() {
		file_proto_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_proto_rawDescData)
	})
	return file_proto_file_proto_rawDescData
}

var file_proto_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_file_proto_goTypes = []interface{}{
	(SaveFileRequest_FileType)(0), // 0: SaveFileRequest.FileType
	(*SaveFileRequest)(nil),       // 1: SaveFileRequest
	(*SaveFileResponse)(nil),      // 2: SaveFileResponse
}
var file_proto_file_proto_depIdxs = []int32{
	0, // 0: SaveFileRequest.fileType:type_name -> SaveFileRequest.FileType
	1, // 1: FileService.SaveFile:input_type -> SaveFileRequest
	2, // 2: FileService.SaveFile:output_type -> SaveFileResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_file_proto_init() }
func file_proto_file_proto_init() {
	if File_proto_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileRequest); i {
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
		file_proto_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveFileResponse); i {
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
			RawDescriptor: file_proto_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_proto_goTypes,
		DependencyIndexes: file_proto_file_proto_depIdxs,
		EnumInfos:         file_proto_file_proto_enumTypes,
		MessageInfos:      file_proto_file_proto_msgTypes,
	}.Build()
	File_proto_file_proto = out.File
	file_proto_file_proto_rawDesc = nil
	file_proto_file_proto_goTypes = nil
	file_proto_file_proto_depIdxs = nil
}
