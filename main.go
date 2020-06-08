package main

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/Mikhalevich/filesharing-file-service/filesystem"
	"github.com/Mikhalevich/filesharing-file-service/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	"github.com/sirupsen/logrus"
)

type params struct {
	serviceName string
	root        string
	temp        string
	permanent   string
	cleanTime   string
}

func loadParams() (*params, error) {
	serviceName := os.Getenv("FS_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "fileservice"
	}

	rootDir := os.Getenv("FS_ROOT_DIR")
	if rootDir == "" {
		rootDir = "storage"
	}

	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	tempDir := os.Getenv("FS_TEMP_DIR")
	if tempDir == "" {
		tempDir = path.Join(os.TempDir(), "Duplo")
	}

	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	permanentDir := os.Getenv("FS_PERMANENT_DIRECTORY")
	if permanentDir == "" {
		permanentDir = "permanent"
	}

	cleanTime := os.Getenv("FS_CLEAN_TIME")
	if cleanTime == "" {
		cleanTime = "23:59"
	}

	return &params{
		serviceName: serviceName,
		root:        rootDir,
		temp:        tempDir,
		permanent:   permanentDir,
		cleanTime:   cleanTime,
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

func makeLoggerWrapper(logger *logrus.Logger) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			logger.Infof("processing %s", req.Method())
			start := time.Now()
			defer logger.Infof("end processing %s, time = %v", req.Method(), time.Now().Sub(start))
			err := fn(ctx, req, rsp)
			if err != nil {
				logger.Errorln(err)
			}
			return err
		}
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

	logger.Infof("running file service with params: service_name = \"%s\" root_direcotry = \"%s\", permanent_directory = \"%s\", cleanTime = \"%s\"", p.serviceName, p.root, p.permanent, p.cleanTime)

	if p.cleanTime != "" {
		err = runCleaner(p.cleanTime, p.root, p.permanent, logger)
		if err != nil {
			logger.Errorln(err)
			return
		}
	}

	service := micro.NewService(
		micro.Name(p.serviceName),
		micro.WrapHandler(makeLoggerWrapper(logger)),
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
}
