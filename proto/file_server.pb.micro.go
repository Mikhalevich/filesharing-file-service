// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: file_server.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FileService service

func NewFileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FileService service

type FileService interface {
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	GetFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (FileService_GetFileService, error)
	UploadFile(ctx context.Context, opts ...client.CallOption) (FileService_UploadFileService, error)
	RemoveFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*EmptyResponse, error)
	IsStorageExists(ctx context.Context, in *IsStorageExistsRequest, opts ...client.CallOption) (*BoolResponse, error)
	CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...client.CallOption) (*CreateStorageResponse, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "FileService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) GetFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (FileService_GetFileService, error) {
	req := c.c.NewRequest(c.name, "FileService.GetFile", &FileRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &fileServiceGetFile{stream}, nil
}

type FileService_GetFileService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Chunk, error)
}

type fileServiceGetFile struct {
	stream client.Stream
}

func (x *fileServiceGetFile) Close() error {
	return x.stream.Close()
}

func (x *fileServiceGetFile) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceGetFile) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceGetFile) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceGetFile) Recv() (*Chunk, error) {
	m := new(Chunk)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileService) UploadFile(ctx context.Context, opts ...client.CallOption) (FileService_UploadFileService, error) {
	req := c.c.NewRequest(c.name, "FileService.UploadFile", &FileUploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &fileServiceUploadFile{stream}, nil
}

type FileService_UploadFileService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*FileUploadRequest) error
	Recv() (*File, error)
}

type fileServiceUploadFile struct {
	stream client.Stream
}

func (x *fileServiceUploadFile) Close() error {
	return x.stream.Close()
}

func (x *fileServiceUploadFile) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceUploadFile) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceUploadFile) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceUploadFile) Send(m *FileUploadRequest) error {
	return x.stream.Send(m)
}

func (x *fileServiceUploadFile) Recv() (*File, error) {
	m := new(File)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileService) RemoveFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "FileService.RemoveFile", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) IsStorageExists(ctx context.Context, in *IsStorageExistsRequest, opts ...client.CallOption) (*BoolResponse, error) {
	req := c.c.NewRequest(c.name, "FileService.IsStorageExists", in)
	out := new(BoolResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...client.CallOption) (*CreateStorageResponse, error) {
	req := c.c.NewRequest(c.name, "FileService.CreateStorage", in)
	out := new(CreateStorageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileService service

type FileServiceHandler interface {
	List(context.Context, *ListRequest, *ListResponse) error
	GetFile(context.Context, *FileRequest, FileService_GetFileStream) error
	UploadFile(context.Context, FileService_UploadFileStream) error
	RemoveFile(context.Context, *FileRequest, *EmptyResponse) error
	IsStorageExists(context.Context, *IsStorageExistsRequest, *BoolResponse) error
	CreateStorage(context.Context, *CreateStorageRequest, *CreateStorageResponse) error
}

func RegisterFileServiceHandler(s server.Server, hdlr FileServiceHandler, opts ...server.HandlerOption) error {
	type fileService interface {
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		GetFile(ctx context.Context, stream server.Stream) error
		UploadFile(ctx context.Context, stream server.Stream) error
		RemoveFile(ctx context.Context, in *FileRequest, out *EmptyResponse) error
		IsStorageExists(ctx context.Context, in *IsStorageExistsRequest, out *BoolResponse) error
		CreateStorage(ctx context.Context, in *CreateStorageRequest, out *CreateStorageResponse) error
	}
	type FileService struct {
		fileService
	}
	h := &fileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FileService{h}, opts...))
}

type fileServiceHandler struct {
	FileServiceHandler
}

func (h *fileServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.FileServiceHandler.List(ctx, in, out)
}

func (h *fileServiceHandler) GetFile(ctx context.Context, stream server.Stream) error {
	m := new(FileRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FileServiceHandler.GetFile(ctx, m, &fileServiceGetFileStream{stream})
}

type FileService_GetFileStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Chunk) error
}

type fileServiceGetFileStream struct {
	stream server.Stream
}

func (x *fileServiceGetFileStream) Close() error {
	return x.stream.Close()
}

func (x *fileServiceGetFileStream) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceGetFileStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceGetFileStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceGetFileStream) Send(m *Chunk) error {
	return x.stream.Send(m)
}

func (h *fileServiceHandler) UploadFile(ctx context.Context, stream server.Stream) error {
	return h.FileServiceHandler.UploadFile(ctx, &fileServiceUploadFileStream{stream})
}

type FileService_UploadFileStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*File) error
	Recv() (*FileUploadRequest, error)
}

type fileServiceUploadFileStream struct {
	stream server.Stream
}

func (x *fileServiceUploadFileStream) Close() error {
	return x.stream.Close()
}

func (x *fileServiceUploadFileStream) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceUploadFileStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceUploadFileStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceUploadFileStream) Send(m *File) error {
	return x.stream.Send(m)
}

func (x *fileServiceUploadFileStream) Recv() (*FileUploadRequest, error) {
	m := new(FileUploadRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *fileServiceHandler) RemoveFile(ctx context.Context, in *FileRequest, out *EmptyResponse) error {
	return h.FileServiceHandler.RemoveFile(ctx, in, out)
}

func (h *fileServiceHandler) IsStorageExists(ctx context.Context, in *IsStorageExistsRequest, out *BoolResponse) error {
	return h.FileServiceHandler.IsStorageExists(ctx, in, out)
}

func (h *fileServiceHandler) CreateStorage(ctx context.Context, in *CreateStorageRequest, out *CreateStorageResponse) error {
	return h.FileServiceHandler.CreateStorage(ctx, in, out)
}
