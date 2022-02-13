package cleaner

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Mikhalevich/filesharing/pkg/service"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	WithField(key string, value interface{}) service.Logger
	WithError(err error) service.Logger
}

type cleaner struct {
	path         string
	protectedDir string
	finish       chan bool
	logger       Logger
}

func New(path string, protectedDirPath string, l Logger) *cleaner {
	path, _ = filepath.Abs(path)

	return &cleaner{
		path:         path,
		protectedDir: protectedDirPath,
		finish:       make(chan bool),
		logger:       l,
	}
}

func (c *cleaner) Run(hour int, minute int) {
	now := time.Now()
	cleanTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	go c.clean(cleanTime)
}

func (c *cleaner) Stop() {
	c.finish <- true
}

func (c *cleaner) clean(t time.Time) {
	tick := func() <-chan time.Time {
		now := time.Now()

		for t.Before(now) {
			t = t.Add(time.Hour * 24)
		}

		return time.After(t.Sub(now))
	}

	for {
		select {
		case <-tick():
		case <-c.finish:
			c.logger.WithField("clean_path", c.path).Info("clean is done")
			return
		}

		storages, err := ioutil.ReadDir(c.path)
		if err != nil {
			log.Println(err)
			c.logger.WithError(err).Error("read dir")
			continue
		}

		for _, storage := range storages {
			if !storage.IsDir() {
				continue
			}

			sPath := path.Join(c.path, storage.Name())

			c.logger.WithField("clean_path", sPath).Info("cleaning")

			files, err := ioutil.ReadDir(sPath)
			if err != nil {
				c.logger.WithError(err).Error("read dir")
				continue
			}

			for _, file := range files {
				if file.IsDir() && file.Name() == c.protectedDir {
					continue
				}

				if err := os.Remove(path.Join(sPath, file.Name())); err != nil {
					c.logger.WithError(err).Error("remove file")
					continue
				}
			}
		}
	}
}
