package filesystem

import (
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
	return f.file.Close()
}

// Name return base file name
func (f *File) Name() string {
	f.file.Close()
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

func (f *File) remove() error {
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
