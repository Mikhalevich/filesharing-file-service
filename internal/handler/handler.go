package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path"

	"github.com/Mikhalevich/filesharing-file-service/internal/filesystem"
	"github.com/Mikhalevich/filesharing/pkg/proto/file"
)

type handler struct {
	storage            *filesystem.Storage
	permanentDirectory string
	tempStorage        *filesystem.Storage
}

func New(s *filesystem.Storage, pd string, temp *filesystem.Storage) *handler {
	return &handler{
		storage:            s,
		permanentDirectory: pd,
		tempStorage:        temp,
	}
}

func (fs *handler) makePath(storage string, isPermanent bool, file string) string {
	if isPermanent {
		return path.Join(storage, fs.permanentDirectory, file)
	}
	return path.Join(storage, file)
}

func wrapError(context, description string, err error) error {
	return fmt.Errorf("[%s] %s: %w", context, description, err)
}

func (fs *handler) List(ctx context.Context, req *file.ListRequest, rsp *file.ListResponse) error {
	fi, err := fs.storage.Files(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""))
	if err != nil {
		return wrapError("List", "get files error", err)
	}

	files := make([]*file.File, 0, len(fi))
	for _, v := range fi {
		files = append(files, &file.File{
			Name:    v.Name(),
			Size:    v.Size(),
			ModTime: v.ModTime().Unix(),
		})
	}

	rsp.Files = files
	return nil
}

func (fs *handler) GetFile(ctx context.Context, req *file.FileRequest, stream file.FileService_GetFileStream) error {
	f, err := fs.storage.File(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return wrapError("GetFile", "get file error", err)
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			stream.Send(&file.Chunk{
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

func contentFromUploadRequest(stream file.FileService_UploadFileStream) ([]byte, error) {
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

func metadataFromUploadRequest(stream file.FileService_UploadFileStream) (*file.FileRequest, error) {
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

func (fs *handler) UploadFile(ctx context.Context, stream file.FileService_UploadFileStream) error {
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

	f, err := fs.tempStorage.Store("", metadata.GetFileName(), r)
	if err != nil {
		return wrapError("UploadFile", "store error", err)
	}

	dir := fs.makePath(metadata.GetStorage(), metadata.GetIsPermanent(), "")
	err = fs.storage.Move(f, dir, metadata.GetFileName())
	if err != nil {
		f.Remove()
		return wrapError("UploadFile", "move error", err)
	}

	err = stream.Send(&file.File{
		Name:    f.Name(),
		Size:    0,
		ModTime: 0,
	})

	if err != nil {
		return wrapError("UploadFile", "send file response error", err)
	}

	return nil
}

func (fs *handler) RemoveFile(ctx context.Context, req *file.FileRequest, rsp *file.EmptyResponse) error {
	err := fs.storage.Remove(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return wrapError("RemoveFile", "remove file error", err)
	}
	return nil
}

func (fs *handler) CreateStorage(ctx context.Context, req *file.CreateStorageRequest, rsp *file.CreateStorageResponse) error {
	status := file.StorageStatus_Ok
	sPath := fs.makePath(req.GetName(), false, "")

	err := fs.storage.Mkdir(sPath)
	if errors.Is(err, filesystem.ErrAlreadyExist) {
		status = file.StorageStatus_AlreadyExist
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

func (fs *handler) IsStorageExists(ctx context.Context, req *file.IsStorageExistsRequest, rsp *file.BoolResponse) error {
	rsp.Flag = fs.storage.IsExist(req.GetName())
	return nil
}
