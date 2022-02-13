package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/asim/go-micro/v3/server"

	"github.com/Mikhalevich/filesharing-file-service/internal/cleaner"
	"github.com/Mikhalevich/filesharing-file-service/internal/filesystem"
	"github.com/Mikhalevich/filesharing-file-service/internal/handler"
	"github.com/Mikhalevich/filesharing/pkg/proto/file"
	"github.com/Mikhalevich/filesharing/pkg/service"
)

type params struct {
	serviceName string
}

type config struct {
	service.Config `yaml:"service"`
	Root           string `yaml:"root_directory"`
	Temp           string `yaml:"temp_directory"`
	Permanent      string `yaml:"permanent_directory"`
	CleanTime      string `yaml:"clean_time"`
}

func (c *config) Service() service.Config {
	return c.Config
}

func (c *config) Validate() error {
	if c.Root == "" {
		c.Root = "storage"
	}

	if c.Temp == "" {
		c.Temp = path.Join(os.TempDir(), "duplo")
	}

	if c.Permanent == "" {
		c.Permanent = "permanent"
	}

	if c.CleanTime == "" {
		c.CleanTime = "23:59"
	}

	return nil
}

func main() {
	var cfg config
	service.Run(os.Getenv("FS_SERVICE_NAME"), &cfg, func(srv server.Server, s service.Servicer) error {
		if err := os.MkdirAll(cfg.Root, os.ModePerm); err != nil {
			return fmt.Errorf("create root directory: %w", err)
		}

		if err := os.MkdirAll(cfg.Temp, os.ModePerm); err != nil {
			return fmt.Errorf("create temp directory: %w", err)
		}

		if cfg.CleanTime != "" {
			t, err := time.Parse("15:04", cfg.CleanTime)
			if err != nil {
				return fmt.Errorf("invalid clean time format: %w", err)
			}

			cleaner := cleaner.New(cfg.Root, cfg.Permanent, s.Logger())
			cleaner.Run(t.Hour(), t.Minute())
		}

		file.RegisterFileServiceHandler(srv, handler.New(
			filesystem.NewStorage(cfg.Root),
			cfg.Permanent,
			filesystem.NewStorage(cfg.Temp),
		))

		return nil
	})
}
