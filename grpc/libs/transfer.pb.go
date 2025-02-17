// protoc --proto_path .\grpc\pb\ --go_out=. --go-grpc_out=. .\grpc\pb\transfer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: transfer.proto

package libs

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

type Ping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ping) Reset() {
	*x = Ping{}
	mi := &file_transfer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Pong struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pong) Reset() {
	*x = Pong{}
	mi := &file_transfer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// tag 标签
// name 文件名
// size 文件大小
// chunks 分片数量
// hash 文件 MD5
type MetaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size          int64                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Chunks        int64                  `protobuf:"varint,4,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Hash          string                 `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaRequest) Reset() {
	*x = MetaRequest{}
	mi := &file_transfer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaRequest) ProtoMessage() {}

func (x *MetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaRequest.ProtoReflect.Descriptor instead.
func (*MetaRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *MetaRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *MetaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MetaRequest) GetChunks() int64 {
	if x != nil {
		return x.Chunks
	}
	return 0
}

func (x *MetaRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// 服务端是否接收数据
type MetaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetaResponse) Reset() {
	*x = MetaResponse{}
	mi := &file_transfer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaResponse) ProtoMessage() {}

func (x *MetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaResponse.ProtoReflect.Descriptor instead.
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *MetaResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *MetaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 文件类型，文件内容和分片
type FileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Chunk         int64                  `protobuf:"varint,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_transfer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *FileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileRequest) GetChunk() int64 {
	if x != nil {
		return x.Chunk
	}
	return 0
}

// 服务端完成接收
type FileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	mi := &file_transfer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *FileResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x40, 0x0a, 0x0c, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x40, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e,
	0x67, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6c, 0x69, 0x62, 0x73, 0x2f, 0x3b, 0x6c, 0x69, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transfer_proto_goTypes = []any{
	(*Ping)(nil),         // 0: pb.Ping
	(*Pong)(nil),         // 1: pb.Pong
	(*MetaRequest)(nil),  // 2: pb.MetaRequest
	(*MetaResponse)(nil), // 3: pb.MetaResponse
	(*FileRequest)(nil),  // 4: pb.FileRequest
	(*FileResponse)(nil), // 5: pb.FileResponse
}
var file_transfer_proto_depIdxs = []int32{
	0, // 0: pb.FileTransferService.ServerCheck:input_type -> pb.Ping
	2, // 1: pb.FileTransferService.SendMeta:input_type -> pb.MetaRequest
	4, // 2: pb.FileTransferService.SendFile:input_type -> pb.FileRequest
	1, // 3: pb.FileTransferService.ServerCheck:output_type -> pb.Pong
	3, // 4: pb.FileTransferService.SendMeta:output_type -> pb.MetaResponse
	5, // 5: pb.FileTransferService.SendFile:output_type -> pb.FileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}
