package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
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

// FileInfo represent file from storage
type FileInfo struct {
	os.FileInfo
}

// Size it's formatted size of the file
// func (fi *FileInfo) Size() string {
// 	return byteSize(fi.FileInfo.Size()).String()
// }

// FileInfoList slice of FileInfo uses for sort
type FileInfoList []FileInfo

func (fil FileInfoList) Len() int {
	return len(fil)
}

func (fil FileInfoList) Swap(i, j int) {
	fil[i], fil[j] = fil[j], fil[i]
}

func (fil FileInfoList) Less(i, j int) bool {
	return fil[i].ModTime().After(fil[j].ModTime())
}

type directory struct {
	Path string
}

func newDirectory(path string) *directory {
	return &directory{
		Path: path,
	}
}

func (d *directory) list() (FileInfoList, error) {
	osFiList, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	fiList := make(FileInfoList, 0, len(osFiList))

	for _, osFi := range osFiList {
		fiList = append(fiList, FileInfo{osFi})
	}

	sort.Sort(fiList)

	return fiList, nil
}

func makeFileName(baseName string, count int) string {
	ext := filepath.Ext(baseName)
	return fmt.Sprintf("%s_%d%s", strings.TrimSuffix(baseName, ext), count, ext)
}

func (d *directory) createUniqueFile(fileName string) (*File, error) {
	for i := 1; i < 100; i++ {
		file, err := createFile(path.Join(d.Path, fileName))
		if os.IsExist(err) {
			fileName = makeFileName(fileName, i)
			continue
		}

		if err != nil {
			return nil, err
		}

		return file, nil
	}

	return nil, nil
}

func (d *directory) moveFile(file *File, fileName string) error {
	for i := 1; i < 100; i++ {
		p := path.Join(d.Path, fileName)
		f, err := openFile(p)
		if os.IsNotExist(err) {
			return file.move(p)
		}

		if err != nil {
			return err
		}
		f.Close()

		fileName = makeFileName(fileName, i)
	}

	return nil
}
