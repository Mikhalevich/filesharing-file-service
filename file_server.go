package main

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/Mikhalevich/file_service/filesystem"
	"github.com/Mikhalevich/file_service/proto"
	"github.com/sirupsen/logrus"
)

type fileServer struct {
	storage            *filesystem.Storage
	permanentDirectory string
	tempStorage        *filesystem.Storage
	logger             *logrus.Logger
}

func newFileServer(s *filesystem.Storage, pd string, temp *filesystem.Storage, l *logrus.Logger) *fileServer {
	return &fileServer{
		storage:            s,
		permanentDirectory: pd,
		tempStorage:        temp,
		logger:             l,
	}
}

func (fs *fileServer) makePath(storage string, isPermanent bool, file string) string {
	if isPermanent {
		return path.Join(storage, fs.permanentDirectory, file)
	}
	return path.Join(storage, file)
}

func (fs *fileServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	fi, err := fs.storage.Files(fs.makePath(req.GetPath(), false, ""))
	if err != nil {
		return nil, err
	}

	files := make([]*proto.File, 0, len(fi))
	for _, v := range fi {
		files = append(files, &proto.File{
			Name:    v.Name(),
			Size:    v.Size(),
			ModTime: v.ModTime().Unix(),
		})
	}

	return &proto.ListResponse{
		Files: files,
	}, nil
}

func (fs *fileServer) GetFile(req *proto.FileRequest, stream proto.FileService_GetFileServer) error {
	file, err := fs.storage.File(fs.makePath(req.GetStorage(), req.GetIsPermanent(), ""), req.GetFileName())
	if err != nil {
		return err
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

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}
	}
	return nil
}

func contentFromUploadRequest(stream proto.FileService_UploadFileServer) ([]byte, error) {
	chunk, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	content := chunk.GetContent()
	if content == nil {
		return nil, errors.New("file content is missing from request")
	}

	return content, nil
}

func metadataFromUploadRequest(stream proto.FileService_UploadFileServer) (*proto.FileRequest, error) {
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

func (fs *fileServer) UploadFile(stream proto.FileService_UploadFileServer) error {
	metadata, err := metadataFromUploadRequest(stream)

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
			w.CloseWithError(err)
			return
		}

		w.Close()
	}()

	file, err := fs.tempStorage.Store("", metadata.GetFileName(), r)

	if err != nil {
		return err
	}

	dir := fs.makePath(metadata.GetStorage(), metadata.GetIsPermanent(), "")
	err = fs.storage.Move(file, dir, metadata.GetFileName())
	if err != nil {
		return err
	}

	stream.SendAndClose(&proto.File{
		Name:    file.Name(),
		Size:    0,
		ModTime: 0,
	})
	return nil
}
