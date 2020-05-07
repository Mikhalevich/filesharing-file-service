package main

import (
	"net"
	"os"
	"path"
	"time"

	"github.com/Mikhalevich/file_service/proto"
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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Errorln(err)
		return
	}

	s := grpc.NewServer()

	proto.RegisterFileServiceServer(s, &fileServer{
		storage:            newFileStorage(p.root),
		permanentDirectory: p.permanent,
		tempStorage:        newFileStorage(p.temp),
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Errorln(err)
		return
	}
}
