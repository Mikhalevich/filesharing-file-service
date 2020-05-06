package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type cleaner struct {
	path         string
	protectedDir string
	finish       chan bool
	logger       *logrus.Logger
}

func newCleaner(path string, protectedDirPath string, l *logrus.Logger) *cleaner {
	path, _ = filepath.Abs(path)

	return &cleaner{
		path:         path,
		protectedDir: protectedDirPath,
		finish:       make(chan bool),
		logger:       l,
	}
}

func (c *cleaner) run(hour int, minute int) {
	now := time.Now()
	cleanTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	go c.clean(cleanTime)
}

func (c *cleaner) stop() {
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
		var now time.Time
		select {
		case now = <-tick():
		case <-c.finish:
			c.logger.Infof("Clean for %s is done\n", c.path)
			return
		}

		c.logger.Infof("time for cleaning: %v", now)

		storages, err := ioutil.ReadDir(c.path)
		if err != nil {
			log.Println(err)
			return
		}

		for _, storage := range storages {
			if !storage.IsDir() {
				continue
			}

			sPath := path.Join(c.path, storage.Name())

			log.Printf("cleaning dir: %q\n", sPath)

			files, err := ioutil.ReadDir(sPath)
			if err != nil {
				log.Println(err)
				return
			}

			for _, file := range files {
				if file.IsDir() && file.Name() == c.protectedDir {
					continue
				}

				err = os.Remove(path.Join(sPath, file.Name()))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
