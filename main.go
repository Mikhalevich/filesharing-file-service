package main

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/Mikhalevich/file_service/filesystem"
	"github.com/Mikhalevich/file_service/proto"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type params struct {
	root      string
	temp      string
	permanent string
	cleanTime string
}

func loadParams() (*params, error) {
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		rootDir = "storage"
	}

	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	tempDir := os.Getenv("TEMP_DIR")
	if tempDir == "" {
		tempDir = path.Join(os.TempDir(), "Duplo")
	}

	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	permanentDir := os.Getenv("PERMANENT_DIRECTORY")
	if permanentDir == "" {
		permanentDir = "permanent"
	}

	cleanTime := os.Getenv("CLEAN_TIME")
	if cleanTime == "" {
		cleanTime = "23:59"
	}

	return &params{
		root:      rootDir,
		temp:      tempDir,
		permanent: permanentDir,
		cleanTime: cleanTime,
	}, nil
}

func runCleaner(cleanTime, rootPath, permanentDirectory string, l *logrus.Logger) error {
	t, err := time.Parse("15:04", cleanTime)
	if err != nil {
		return err
	}

	cleaner := newCleaner(rootPath, permanentDirectory, l)
	cleaner.run(t.Hour(), t.Minute())

	return nil
}

func makeUnaryInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		h, err := handler(ctx, req)
		if err != nil {
			logger.Errorln(err)
		}
		return h, err
	}
}

func makeStreamInterceptor(logger *logrus.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		if err != nil {
			logger.Errorln(err)
		}
		return err
	}
}

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	p, err := loadParams()
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Infof("running file service with params: root_direcotry = \"%s\", permanent_directory = \"%s\", cleanTime = \"%s\"", p.root, p.permanent, p.cleanTime)

	if p.cleanTime != "" {
		err = runCleaner(p.cleanTime, p.root, p.permanent, logger)
		if err != nil {
			logger.Errorln(err)
			return
		}
	}

	service := micro.NewService(
		micro.Name("fileservice"),
	)

	service.Init()

	proto.RegisterFileServiceHandler(service.Server(), &fileServer{
		storage:            filesystem.NewStorage(p.root),
		permanentDirectory: p.permanent,
		tempStorage:        filesystem.NewStorage(p.temp),
	})

	err = service.Run()
	if err != nil {
		logger.Errorln(err)
		return
	}

	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	logger.Errorln(err)
	// 	return
	// }

	// s := grpc.NewServer(grpc.UnaryInterceptor(makeUnaryInterceptor(logger)),
	// 	grpc.StreamInterceptor(makeStreamInterceptor(logger)))

	// proto.RegisterFileServiceServer(s, &fileServer{
	// 	storage:            filesystem.NewStorage(p.root),
	// 	permanentDirectory: p.permanent,
	// 	tempStorage:        filesystem.NewStorage(p.temp),
	// })

	// err = s.Serve(lis)
	// if err != nil {
	// 	logger.Errorln(err)
	// 	return
	// }
}
