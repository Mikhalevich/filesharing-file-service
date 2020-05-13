package filesystem

import (
	"fmt"
	"io"
	"path"
)

// Storage represent file manipulation fasade
type Storage struct {
	rootPath string
}

// NewStorage create a new Storage object with abs root path
func NewStorage(root string) *Storage {
	return &Storage{
		rootPath: root,
	}
}

func (s *Storage) join(p string) string {
	return path.Join(s.rootPath, p)
}

// Files returns files from current storage directory
func (s *Storage) Files(p string) (FileInfoList, error) {
	l, err := newDirectory(s.join(p)).list()
	if err != nil {
		return nil, fmt.Errorf("[files] unable to get list of files, err = %w", err)
	}

	return l, nil
}

// File return existing file with name fileName insice directory dir
func (s *Storage) File(dir, fileName string) (*File, error) {
	file, err := openFile(s.join(path.Join(dir, fileName)))
	if err != nil {
		err = fmt.Errorf("[get file] dir = %s, name = %s, err = %w", dir, fileName, err)
	}
	return file, err
}

// Store save file with fileName inside direcory dir, returns a closed newly created file object
func (s *Storage) Store(dir string, fileName string, data io.Reader) (*File, error) {
	f, err := newDirectory(s.join(dir)).createUniqueFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("[store file] unable to create file dir = %s, name = %s, err = %w", dir, fileName, err)
	}
	defer f.Close()

	_, err = io.Copy(f, data)
	if err != nil {
		return nil, fmt.Errorf("[store file] error while copy file data dir = %s, name = %s, err = %w", dir, fileName, err)
	}

	return f, nil
}

// Move move file into direcotyr dir with name fileName
func (s *Storage) Move(file *File, dir string, fileName string) error {
	err := newDirectory(s.join(dir)).moveFile(file, fileName)
	if err != nil {
		return fmt.Errorf("[move file] unable to move file dir = %s, name = %s, err = %w", dir, fileName, err)
	}
	return nil
}

// Remove just remove file from storage
func (s *Storage) Remove(dir string, fileName string) error {
	f, err := openFile(s.join(path.Join(dir, fileName)))
	if err != nil {
		return fmt.Errorf("[remove file] unable to open file dir = %s, name = %s, err = %w", dir, fileName, err)
	}
	err = f.remove()
	if err != nil {
		return fmt.Errorf("[remove file] unable to remove file dir = %s, name = %s, err = %w", dir, fileName, err)
	}
	return nil
}
