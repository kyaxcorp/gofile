package filesystem

import (
	"io/ioutil"
)

//type File fs.File

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	Path string
}

func (f *File) Read(b []byte) (int, error) {
	b, _err := ioutil.ReadFile(f.Path)
	if _err != nil {
		return 0, _err
	}
	return len(b), nil
}

func (f *File) Delete() error {
	return nil
}
