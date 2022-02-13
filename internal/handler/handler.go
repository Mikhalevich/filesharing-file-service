package handler

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/Mikhalevich/filesharing-file-service/internal/filesystem"
	"github.com/Mikhalevich/filesharing/pkg/httperror"
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

func (fs *handler) List(ctx context.Context, req *file.ListRequest, rsp *file.ListResponse) error {
	fi, err := fs.storage.Files(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""))
	if err != nil {
		return httperror.NewInternalError("get files error").WithError(err)
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
		return httperror.NewInternalError("get file error").WithError(err)
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

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return httperror.NewInternalError("send data error").WithError(err)
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
		return httperror.NewInternalError("metadate error").WithError(err)
	}

	r, w := io.Pipe()

	go func() {
		var closeError error
		for {
			content, err := contentFromUploadRequest(stream)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				closeError = err
				break
			}
			w.Write(content)
		}

		if closeError != nil {
			w.CloseWithError(httperror.NewInternalError("close file error").WithError(closeError))
			return
		}

		w.Close()
	}()

	f, err := fs.tempStorage.Store("", metadata.GetFileName(), r)
	if err != nil {
		return httperror.NewInternalError("store error").WithError(err)
	}

	dir := fs.makePath(metadata.GetStorage(), metadata.GetIsPermanent(), "")
	if err := fs.storage.Move(f, dir, metadata.GetFileName()); err != nil {
		f.Remove()
		return httperror.NewInternalError("move error").WithError(err)
	}

	if err := stream.Send(&file.File{
		Name:    f.Name(),
		Size:    0,
		ModTime: 0,
	}); err != nil {
		return httperror.NewInternalError("send file response error").WithError(err)
	}

	return nil
}

func (fs *handler) RemoveFile(ctx context.Context, req *file.FileRequest, rsp *file.RemoveFileResponse) error {
	err := fs.storage.Remove(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return httperror.NewInternalError("remove file error").WithError(err)
	}
	return nil
}

func (fs *handler) CreateStorage(ctx context.Context, req *file.CreateStorageRequest, rsp *file.CreateStorageResponse) error {
	sPath := fs.makePath(req.GetName(), false, "")

	if err := fs.storage.Mkdir(sPath); err != nil {
		if errors.Is(err, filesystem.ErrAlreadyExist) {
			return httperror.NewAlreadyExistError("storage already exist")
		}
		return httperror.NewInternalError("create folder error").WithError(err)
	}

	if req.GetWithPermanent() {
		if err := fs.storage.Mkdir(fs.makePath(req.GetName(), true, "")); err != nil {
			if !errors.Is(err, filesystem.ErrAlreadyExist) {
				fs.storage.RemoveDir(sPath)
				return httperror.NewInternalError("create permanent folder error").WithError(err)
			}
		}
	}

	return nil
}

func (fs *handler) IsStorageExists(ctx context.Context, req *file.IsStorageExistsRequest, rsp *file.BoolResponse) error {
	rsp.Flag = fs.storage.IsExist(req.GetName())
	return nil
}
