package main

import (
	"context"

	"github.com/Mikhalevich/file_service/proto"
)

type fileServer struct {
	storage     *fileStorage
	tempStorage *fileStorage
}

func newFileServer(s *fileStorage, temp *fileStorage) *fileServer {
	return &fileServer{
		storage:     s,
		tempStorage: temp,
	}
}

func (fs *fileServer) Files(ctx context.Context, req *proto.FilesRequest) (*proto.FilesResponse, error) {
	fi := fs.storage.files(req.GetPath())
	files := make([]*proto.File, 0, len(fi))
	for _, v := range fi {
		files = append(files, &proto.File{
			Name:    v.Name(),
			Size:    v.Size(),
			ModTime: v.ModTime().Unix(),
		})
	}

	return &proto.FilesResponse{
		Files: files,
	}, nil
}
