package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path"

	"github.com/Mikhalevich/file_service/filesystem"
	"github.com/Mikhalevich/file_service/proto"
)

type fileServer struct {
	storage            *filesystem.Storage
	permanentDirectory string
	tempStorage        *filesystem.Storage
}

func newFileServer(s *filesystem.Storage, pd string, temp *filesystem.Storage) *fileServer {
	return &fileServer{
		storage:            s,
		permanentDirectory: pd,
		tempStorage:        temp,
	}
}

func (fs *fileServer) makePath(storage string, isPermanent bool, file string) string {
	if isPermanent {
		return path.Join(storage, fs.permanentDirectory, file)
	}
	return path.Join(storage, file)
}

func wrapError(context, description string, err error) error {
	return fmt.Errorf("[%s] %s: %w", context, description, err)
}

func (fs *fileServer) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	fi, err := fs.storage.Files(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""))
	if err != nil {
		return wrapError("List", "get files error", err)
	}

	files := make([]*proto.File, 0, len(fi))
	for _, v := range fi {
		files = append(files, &proto.File{
			Name:    v.Name(),
			Size:    v.Size(),
			ModTime: v.ModTime().Unix(),
		})
	}

	rsp.Files = files
	return nil
}

func (fs *fileServer) GetFile(ctx context.Context, req *proto.FileRequest, stream proto.FileService_GetFileStream) error {
	file, err := fs.storage.File(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return wrapError("GetFile", "get file error", err)
	}
	defer file.Close()

	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			stream.Send(&proto.Chunk{
				Content: buf[:n],
			})
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return wrapError("GetFile", "send data error", err)
		}
	}
	return nil
}

func contentFromUploadRequest(stream proto.FileService_UploadFileStream) ([]byte, error) {
	chunk, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	if chunk.GetEnd() {
		return nil, io.EOF
	}

	content := chunk.GetContent()
	if content == nil {
		return nil, errors.New("file content is missing from request")
	}

	return content, nil
}

func metadataFromUploadRequest(stream proto.FileService_UploadFileStream) (*proto.FileRequest, error) {
	chunk, err := stream.Recv()
	if err != nil {
		return nil, err

	}
	metadata := chunk.GetMetadata()
	if metadata == nil {
		return nil, errors.New("file info is missing from request")
	}

	return metadata, nil
}

func (fs *fileServer) UploadFile(ctx context.Context, stream proto.FileService_UploadFileStream) error {
	metadata, err := metadataFromUploadRequest(stream)
	if err != nil {
		return wrapError("UploadFile", "metadata error", err)
	}

	r, w := io.Pipe()

	go func() {
		var closeError error
		for {
			content, err := contentFromUploadRequest(stream)
			if err == io.EOF {
				break
			} else if err != nil {
				closeError = err
				break
			}
			w.Write(content)
		}

		if closeError != nil {
			w.CloseWithError(wrapError("UploadFile", "write error", closeError))
			return
		}

		w.Close()
	}()

	file, err := fs.tempStorage.Store("", metadata.GetFileName(), r)
	if err != nil {
		return wrapError("UploadFile", "store error", err)
	}

	dir := fs.makePath(metadata.GetStorage(), metadata.GetIsPermanent(), "")
	err = fs.storage.Move(file, dir, metadata.GetFileName())
	if err != nil {
		file.Remove()
		return wrapError("UploadFile", "move error", err)
	}

	err = stream.Send(&proto.File{
		Name:    file.Name(),
		Size:    0,
		ModTime: 0,
	})

	if err != nil {
		return wrapError("UploadFile", "send file response error", err)
	}

	return nil
}

func (fs *fileServer) RemoveFile(ctx context.Context, req *proto.FileRequest, rsp *proto.EmptyResponse) error {
	err := fs.storage.Remove(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return wrapError("RemoveFile", "remove file error", err)
	}
	return nil
}

func (fs *fileServer) CreateStorage(ctx context.Context, req *proto.CreateStorageRequest, rsp *proto.CreateStorageResponse) error {
	status := proto.StorageStatus_Ok
	sPath := fs.makePath(req.GetName(), false, "")

	err := fs.storage.Mkdir(sPath)
	if errors.Is(err, filesystem.ErrAlreadyExist) {
		status = proto.StorageStatus_AlreadyExist
	} else if err != nil {
		return wrapError("CreateStorage", "create folder error", err)
	}

	if req.GetWithPermanent() {
		err = fs.storage.Mkdir(fs.makePath(req.GetName(), true, ""))
		if (err != nil) && !errors.Is(err, filesystem.ErrAlreadyExist) {
			fs.storage.RemoveDir(sPath)
			return wrapError("CreateStorage", "create permanent folder error", err)
		}
	}

	rsp.Status = status
	return nil
}

func (fs *fileServer) IsStorageExists(ctx context.Context, req *proto.IsStorageExistsRequest, rsp *proto.BoolResponse) error {
	rsp.Flag = fs.storage.IsExist(req.GetName())
	return nil
}
