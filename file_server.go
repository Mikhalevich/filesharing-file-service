package main

import (
	"context"
	"errors"
	"io"
	"os"
	"path"

	"github.com/Mikhalevich/file_service/proto"
)

type fileServer struct {
	storage            *fileStorage
	permanentDirectory string
	tempStorage        *fileStorage
}

func newFileServer(s *fileStorage, pd string, temp *fileStorage) *fileServer {
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

func (fs *fileServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	fi := fs.storage.files(fs.makePath(req.GetPath(), false, ""))
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
	p := fs.makePath(req.GetStorage(), req.GetIsPermanent(), req.GetFileName())
	if !fs.storage.isExists(p) {
		return errors.New("file doesn't exist")
	}

	file, err := os.Open(fs.storage.join(p))
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
