package filesystem

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

type byteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kb byteSize = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

var (
	permanentDir string

	errNotExists = errors.New("Not exists")
)

func (b byteSize) String() string {
	switch {
	case b >= yb:
		return fmt.Sprintf("%.2fYB", b/yb)
	case b >= zb:
		return fmt.Sprintf("%.2fZB", b/zb)
	case b >= eb:
		return fmt.Sprintf("%.2fEB", b/eb)
	case b >= pb:
		return fmt.Sprintf("%.2fPB", b/pb)
	case b >= tb:
		return fmt.Sprintf("%.2fTB", b/tb)
	case b >= gb:
		return fmt.Sprintf("%.2fGB", b/gb)
	case b >= mb:
		return fmt.Sprintf("%.2fMB", b/mb)
	case b >= kb:
		return fmt.Sprintf("%.2fKB", b/kb)
	}
	return fmt.Sprintf("%.2fB", b)
}

type fileInfo struct {
	os.FileInfo
}

func (fi *fileInfo) size() string {
	return byteSize(fi.FileInfo.Size()).String()
}

type fileInfoList []fileInfo

func (fil fileInfoList) Len() int {
	return len(fil)
}

func (fil fileInfoList) Swap(i, j int) {
	fil[i], fil[j] = fil[j], fil[i]
}

func (fil fileInfoList) Less(i, j int) bool {
	if permanentDir != "" {
		if fil[i].IsDir() && fil[i].Name() == permanentDir {
			return true
		}

		if fil[j].IsDir() && fil[j].Name() == permanentDir {
			return false
		}
	}

	return fil[i].ModTime().After(fil[j].ModTime())
}

func (fil fileInfoList) Exist(name string) bool {
	for _, fi := range fil {
		if fi.Name() == name {
			return true
		}
	}

	return false
}

type Storage struct {
	rootPath string
}

func NewStorage(root string) *Storage {
	return &Storage{
		rootPath: root,
	}
}

func (s *Storage) root() string {
	return s.rootPath
}

func (s *Storage) Join(p string) string {
	return path.Join(s.rootPath, p)
}

func (s *Storage) IsExists(p string) bool {
	_, err := os.Stat(s.Join(p))
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func (s *Storage) mkdir(dir string) error {
	return os.Mkdir(s.Join(dir), os.ModePerm)
}

func (s *Storage) Files(p string) fileInfoList {
	return newDirectory(s.Join(p)).list()
}

func (s *Storage) Store(dir string, fileName string, data io.Reader) (string, error) {
	dirPath := s.Join(dir)
	uniqueName := newDirectory(dirPath).uniqueName(fileName)
	f, err := os.Create(path.Join(dirPath, uniqueName))
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, data)
	if err != nil {
		return "", err
	}

	return uniqueName, nil
}

func (s *Storage) Move(filePath string, dir string, fileName string) error {
	dirPath := s.Join(dir)
	uniqueName := newDirectory(dirPath).uniqueName(fileName)
	return os.Rename(filePath, path.Join(dirPath, uniqueName))
}

func (s *Storage) Remove(dir string, fileName string) error {
	dirPath := s.Join(dir)
	files := newDirectory(dirPath).list()
	if !files.Exist(fileName) {
		return errNotExists
	}

	return os.Remove(path.Join(dirPath, fileName))
}
