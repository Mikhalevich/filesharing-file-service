package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type params struct {
	root      string
	permanent string
	cleanTime string
}

func loadParams() *params {
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		rootDir = "storage"
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
		permanent: permanentDir,
		cleanTime: cleanTime,
	}
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

	p := loadParams()

	logger.Infof("running file service with params: root_direcotry = \"%s\", permanent_directory = \"%s\", cleanTime = \"%s\"", p.root, p.permanent, p.cleanTime)

	if p.cleanTime != "" {
		err := runCleaner(p.cleanTime, p.root, p.permanent, logger)
		if err != nil {
			logger.Errorln(err)
			return
		}
	}
}
