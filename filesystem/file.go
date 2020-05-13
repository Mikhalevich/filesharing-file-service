package filesystem

import (
	"errors"
	"os"
)

// File represent single file inside file storage
type File struct {
	absPath string
	file    *os.File
}

func createFile(p string) (*File, error) {
	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return &File{
		absPath: p,
		file:    f,
	}, nil
}

func openFile(p string) (*File, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	return &File{
		absPath: p,
		file:    f,
	}, nil
}

// Close open file
func (f *File) Close() error {
	err := f.file.Close()
	if errors.Is(err, os.ErrClosed) {
		return nil
	}
	return err
}

// Name return base file name
func (f *File) Name() string {
	fi, err := os.Stat(f.absPath)
	if err != nil {
		return ""
	}
	return fi.Name()
}

func (f *File) Read(p []byte) (n int, err error) {
	return f.file.Read(p)
}

func (f *File) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}

// Remove deletes file from filesystem
func (f *File) Remove() error {
	f.Close()
	return os.Remove(f.absPath)
}

func (f *File) move(p string) error {
	f.Close()
	err := os.Rename(f.absPath, p)
	if err != nil {
		return err
	}

	f.absPath = p
	return nil
}
