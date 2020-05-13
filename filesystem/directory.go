package filesystem

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type directory struct {
	Path string
}

func newDirectory(path string) *directory {
	return &directory{
		Path: path,
	}
}

func (d *directory) list() fileInfoList {
	osFiList, err := ioutil.ReadDir(d.Path)
	if err != nil {
		log.Println(err)
		return fileInfoList{}
	}

	fiList := make(fileInfoList, 0, len(osFiList))

	for _, osFi := range osFiList {
		fiList = append(fiList, fileInfo{osFi})
	}

	sort.Sort(fiList)

	return fiList
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

// func (d *directory) uniqueName(fileName string) string {
// 	ld := d.list()
// 	if !ld.Exist(fileName) {
// 		return fileName
// 	}

// 	ext := filepath.Ext(fileName)

// 	nameTemplate := fmt.Sprintf("%s%s%s", strings.TrimSuffix(fileName, ext), "_%d", ext)

// 	for count := 1; ; count++ {
// 		fileName = fmt.Sprintf(nameTemplate, count)
// 		if !ld.Exist(fileName) {
// 			break
// 		}
// 	}

// 	return fileName
// }
