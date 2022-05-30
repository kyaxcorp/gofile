package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
	"io/ioutil"
)

//type File fs.File

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	FPath string

	info FileInfo
}

func newFile(filePath string) (*File, error) {
	// Generate file info
	fInfo, _err := newFileInfo(filePath)
	if _err != nil {
		return nil, _err
	}

	// generate the file itself
	f := &File{
		FPath: filePath,
		info:  *fInfo,
	}

	// return
	return f, nil
}

func (f *File) Read(b []byte) (int, error) {
	b, _err := ioutil.ReadFile(f.FPath)
	if _err != nil {
		return 0, _err
	}
	return len(b), nil
}

func (f *File) Delete() error {
	return nil
}

func (f *File) Info() driver.FileInfoInterface {
	return &f.info
}
