// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file_server.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StorageStatus int32

const (
	StorageStatus_Ok           StorageStatus = 0
	StorageStatus_AlreadyExist StorageStatus = 1
)

var StorageStatus_name = map[int32]string{
	0: "Ok",
	1: "AlreadyExist",
}

var StorageStatus_value = map[string]int32{
	"Ok":           0,
	"AlreadyExist": 1,
}

func (x StorageStatus) String() string {
	return proto.EnumName(StorageStatus_name, int32(x))
}

func (StorageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{0}
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{0}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type ListRequest struct {
	Storage              string   `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	IsPermanent          bool     `protobuf:"varint,2,opt,name=isPermanent,proto3" json:"isPermanent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *ListRequest) GetIsPermanent() bool {
	if m != nil {
		return m.IsPermanent
	}
	return false
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	ModTime              int64    `protobuf:"varint,3,opt,name=modTime,proto3" json:"modTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{2}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetModTime() int64 {
	if m != nil {
		return m.ModTime
	}
	return 0
}

type ListResponse struct {
	Files                []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type FileRequest struct {
	Storage              string   `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	IsPermanent          bool     `protobuf:"varint,2,opt,name=isPermanent,proto3" json:"isPermanent,omitempty"`
	FileName             string   `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{4}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *FileRequest) GetIsPermanent() bool {
	if m != nil {
		return m.IsPermanent
	}
	return false
}

func (m *FileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{5}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FileUploadRequest struct {
	// Types that are valid to be assigned to FileChunk:
	//	*FileUploadRequest_Metadata
	//	*FileUploadRequest_Content
	//	*FileUploadRequest_End
	FileChunk            isFileUploadRequest_FileChunk `protobuf_oneof:"fileChunk"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FileUploadRequest) Reset()         { *m = FileUploadRequest{} }
func (m *FileUploadRequest) String() string { return proto.CompactTextString(m) }
func (*FileUploadRequest) ProtoMessage()    {}
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{6}
}

func (m *FileUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadRequest.Unmarshal(m, b)
}
func (m *FileUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadRequest.Marshal(b, m, deterministic)
}
func (m *FileUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadRequest.Merge(m, src)
}
func (m *FileUploadRequest) XXX_Size() int {
	return xxx_messageInfo_FileUploadRequest.Size(m)
}
func (m *FileUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadRequest proto.InternalMessageInfo

type isFileUploadRequest_FileChunk interface {
	isFileUploadRequest_FileChunk()
}

type FileUploadRequest_Metadata struct {
	Metadata *FileRequest `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type FileUploadRequest_Content struct {
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

type FileUploadRequest_End struct {
	End bool `protobuf:"varint,3,opt,name=end,proto3,oneof"`
}

func (*FileUploadRequest_Metadata) isFileUploadRequest_FileChunk() {}

func (*FileUploadRequest_Content) isFileUploadRequest_FileChunk() {}

func (*FileUploadRequest_End) isFileUploadRequest_FileChunk() {}

func (m *FileUploadRequest) GetFileChunk() isFileUploadRequest_FileChunk {
	if m != nil {
		return m.FileChunk
	}
	return nil
}

func (m *FileUploadRequest) GetMetadata() *FileRequest {
	if x, ok := m.GetFileChunk().(*FileUploadRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *FileUploadRequest) GetContent() []byte {
	if x, ok := m.GetFileChunk().(*FileUploadRequest_Content); ok {
		return x.Content
	}
	return nil
}

func (m *FileUploadRequest) GetEnd() bool {
	if x, ok := m.GetFileChunk().(*FileUploadRequest_End); ok {
		return x.End
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileUploadRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileUploadRequest_Metadata)(nil),
		(*FileUploadRequest_Content)(nil),
		(*FileUploadRequest_End)(nil),
	}
}

type IsStorageExistsRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsStorageExistsRequest) Reset()         { *m = IsStorageExistsRequest{} }
func (m *IsStorageExistsRequest) String() string { return proto.CompactTextString(m) }
func (*IsStorageExistsRequest) ProtoMessage()    {}
func (*IsStorageExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{7}
}

func (m *IsStorageExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsStorageExistsRequest.Unmarshal(m, b)
}
func (m *IsStorageExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsStorageExistsRequest.Marshal(b, m, deterministic)
}
func (m *IsStorageExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsStorageExistsRequest.Merge(m, src)
}
func (m *IsStorageExistsRequest) XXX_Size() int {
	return xxx_messageInfo_IsStorageExistsRequest.Size(m)
}
func (m *IsStorageExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsStorageExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsStorageExistsRequest proto.InternalMessageInfo

func (m *IsStorageExistsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BoolResponse struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolResponse) Reset()         { *m = BoolResponse{} }
func (m *BoolResponse) String() string { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()    {}
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{8}
}

func (m *BoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolResponse.Unmarshal(m, b)
}
func (m *BoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolResponse.Marshal(b, m, deterministic)
}
func (m *BoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolResponse.Merge(m, src)
}
func (m *BoolResponse) XXX_Size() int {
	return xxx_messageInfo_BoolResponse.Size(m)
}
func (m *BoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BoolResponse proto.InternalMessageInfo

func (m *BoolResponse) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

type CreateStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	WithPermanent        bool     `protobuf:"varint,2,opt,name=withPermanent,proto3" json:"withPermanent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStorageRequest) Reset()         { *m = CreateStorageRequest{} }
func (m *CreateStorageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStorageRequest) ProtoMessage()    {}
func (*CreateStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{9}
}

func (m *CreateStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStorageRequest.Unmarshal(m, b)
}
func (m *CreateStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStorageRequest.Marshal(b, m, deterministic)
}
func (m *CreateStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStorageRequest.Merge(m, src)
}
func (m *CreateStorageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStorageRequest.Size(m)
}
func (m *CreateStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStorageRequest proto.InternalMessageInfo

func (m *CreateStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateStorageRequest) GetWithPermanent() bool {
	if m != nil {
		return m.WithPermanent
	}
	return false
}

type CreateStorageResponse struct {
	Status               StorageStatus `protobuf:"varint,1,opt,name=status,proto3,enum=StorageStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateStorageResponse) Reset()         { *m = CreateStorageResponse{} }
func (m *CreateStorageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStorageResponse) ProtoMessage()    {}
func (*CreateStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a16f8fdec850da66, []int{10}
}

func (m *CreateStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStorageResponse.Unmarshal(m, b)
}
func (m *CreateStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStorageResponse.Marshal(b, m, deterministic)
}
func (m *CreateStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStorageResponse.Merge(m, src)
}
func (m *CreateStorageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStorageResponse.Size(m)
}
func (m *CreateStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStorageResponse proto.InternalMessageInfo

func (m *CreateStorageResponse) GetStatus() StorageStatus {
	if m != nil {
		return m.Status
	}
	return StorageStatus_Ok
}

func init() {
	proto.RegisterEnum("StorageStatus", StorageStatus_name, StorageStatus_value)
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
	proto.RegisterType((*FileRequest)(nil), "FileRequest")
	proto.RegisterType((*Chunk)(nil), "Chunk")
	proto.RegisterType((*FileUploadRequest)(nil), "FileUploadRequest")
	proto.RegisterType((*IsStorageExistsRequest)(nil), "IsStorageExistsRequest")
	proto.RegisterType((*BoolResponse)(nil), "BoolResponse")
	proto.RegisterType((*CreateStorageRequest)(nil), "CreateStorageRequest")
	proto.RegisterType((*CreateStorageResponse)(nil), "CreateStorageResponse")
}

func init() { proto.RegisterFile("file_server.proto", fileDescriptor_a16f8fdec850da66) }

var fileDescriptor_a16f8fdec850da66 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xb5, 0x9b, 0x34, 0x1f, 0x63, 0x3b, 0x4d, 0x47, 0xb4, 0x44, 0xe6, 0x12, 0x96, 0x0f, 0x85,
	0x52, 0x56, 0x55, 0x38, 0xf6, 0x00, 0xa4, 0x2a, 0xa4, 0x12, 0x82, 0x6a, 0x03, 0x17, 0x2e, 0x68,
	0x69, 0xa6, 0xad, 0x55, 0x7f, 0x04, 0xef, 0x26, 0x50, 0xc4, 0xdf, 0xe2, 0xff, 0xa1, 0x5d, 0xc7,
	0x21, 0x09, 0x11, 0x17, 0x4e, 0xde, 0x99, 0x1d, 0xbf, 0x79, 0x33, 0xef, 0xd9, 0xb0, 0x7b, 0x19,
	0xc5, 0xf4, 0x59, 0x51, 0x3e, 0xa3, 0x9c, 0x4f, 0xf2, 0x4c, 0x67, 0x6c, 0x07, 0x82, 0xd3, 0x64,
	0xa2, 0x6f, 0x05, 0xa9, 0x49, 0x96, 0x2a, 0x62, 0x67, 0xe0, 0xbd, 0x8d, 0x94, 0x16, 0xf4, 0x75,
	0x4a, 0x4a, 0x63, 0x07, 0xea, 0x4a, 0x67, 0xb9, 0xbc, 0xa2, 0x8e, 0xdb, 0x75, 0x7b, 0x4d, 0x51,
	0x86, 0xd8, 0x05, 0x2f, 0x52, 0xe7, 0x94, 0x27, 0x32, 0xa5, 0x54, 0x77, 0xb6, 0xba, 0x6e, 0xaf,
	0x21, 0x96, 0x53, 0x6c, 0x08, 0xd5, 0xd7, 0x51, 0x4c, 0x88, 0x50, 0x4d, 0x65, 0x52, 0x02, 0xd8,
	0xb3, 0xc9, 0xa9, 0xe8, 0x07, 0xd9, 0xd7, 0x2a, 0xc2, 0x9e, 0x4d, 0xaf, 0x24, 0x1b, 0x7f, 0x88,
	0x12, 0xea, 0x54, 0x6c, 0xba, 0x0c, 0xd9, 0x53, 0xf0, 0x0b, 0x52, 0x05, 0x49, 0xbc, 0x07, 0xdb,
	0x66, 0x14, 0xd5, 0x71, 0xbb, 0x95, 0x9e, 0xd7, 0xdf, 0xe6, 0xa6, 0x8f, 0x28, 0x72, 0x8c, 0xc0,
	0xb3, 0xe1, 0xff, 0x4f, 0x80, 0x21, 0x34, 0x0c, 0xe6, 0x3b, 0x39, 0xa7, 0xd4, 0x14, 0x8b, 0x98,
	0xdd, 0x87, 0xed, 0x93, 0xeb, 0x69, 0x7a, 0x63, 0x1a, 0x5c, 0x64, 0xa9, 0x36, 0x10, 0xa6, 0x81,
	0x2f, 0xca, 0x90, 0xfd, 0x84, 0x5d, 0xc3, 0xe4, 0xe3, 0x24, 0xce, 0xe4, 0xb8, 0xe4, 0x73, 0x00,
	0x8d, 0x84, 0xb4, 0x1c, 0x4b, 0x2d, 0x6d, 0xbd, 0xd7, 0xf7, 0xf9, 0x12, 0xdf, 0xa1, 0x23, 0x16,
	0xf7, 0x18, 0xfe, 0x81, 0x36, 0xec, 0xfc, 0xa1, 0xb3, 0x00, 0x47, 0x84, 0x0a, 0xa5, 0x63, 0x4b,
	0xab, 0x31, 0x74, 0x84, 0x09, 0x06, 0x1e, 0x34, 0x0d, 0x3f, 0xcb, 0x8b, 0x1d, 0xc2, 0xfe, 0x99,
	0x1a, 0x15, 0xb3, 0x9e, 0x7e, 0x8f, 0x94, 0x56, 0x25, 0x85, 0x0d, 0x82, 0x30, 0x06, 0xfe, 0x20,
	0xcb, 0xe2, 0xc5, 0x8a, 0x11, 0xaa, 0x97, 0xb1, 0xbc, 0xb2, 0x35, 0x0d, 0x61, 0xcf, 0xec, 0x1c,
	0xee, 0x9c, 0xe4, 0x24, 0x35, 0xcd, 0x51, 0xff, 0x81, 0x87, 0x0f, 0x21, 0xf8, 0x16, 0xe9, 0xeb,
	0xf5, 0xf5, 0xae, 0x26, 0xd9, 0x0b, 0xd8, 0x5b, 0x43, 0x9c, 0xb7, 0x7f, 0x0c, 0x35, 0xa5, 0xa5,
	0x9e, 0x2a, 0x0b, 0xda, 0xea, 0xb7, 0xf8, 0xbc, 0x62, 0x64, 0xb3, 0x62, 0x7e, 0x7b, 0xf0, 0x04,
	0x82, 0x95, 0x0b, 0xac, 0xc1, 0xd6, 0xfb, 0x9b, 0xb6, 0x83, 0x6d, 0xf0, 0x5f, 0xc5, 0x39, 0xc9,
	0xf1, 0xad, 0x9d, 0xbd, 0xed, 0xf6, 0x7f, 0x6d, 0x15, 0xc6, 0x18, 0x51, 0x3e, 0x8b, 0x2e, 0x08,
	0x1f, 0x41, 0xd5, 0x98, 0x0a, 0x7d, 0xbe, 0x64, 0xf8, 0x30, 0xe0, 0xcb, 0x4e, 0x63, 0x0e, 0x3e,
	0x80, 0xfa, 0x1b, 0xd2, 0xd6, 0xc8, 0x2b, 0x42, 0x85, 0x35, 0x5e, 0xec, 0xd9, 0x39, 0x72, 0xf1,
	0x19, 0x40, 0xa1, 0x72, 0x61, 0x78, 0xfe, 0x97, 0xec, 0x61, 0xe1, 0x51, 0xe6, 0xf4, 0xdc, 0x23,
	0x17, 0x0f, 0x01, 0x04, 0x25, 0xd9, 0x8c, 0x36, 0xc0, 0xb6, 0xf8, 0xea, 0x07, 0xe9, 0xe0, 0x31,
	0xec, 0xac, 0x09, 0x89, 0x77, 0xf9, 0x66, 0x69, 0xc3, 0x80, 0x2f, 0xab, 0xc8, 0x1c, 0x7c, 0x09,
	0xc1, 0xca, 0x86, 0x71, 0x8f, 0x6f, 0xd2, 0x30, 0xdc, 0xe7, 0x1b, 0x85, 0x60, 0xce, 0xa0, 0xf9,
	0xa9, 0xce, 0x8f, 0xed, 0xdf, 0xe2, 0x4b, 0xcd, 0x3e, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0x7e, 0x48, 0x0a, 0x49, 0x04, 0x00, 0x00,
}
