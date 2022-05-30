package filesystem

import "github.com/kyaxcorp/gofile"

func (f *File) Info() gofile.FileInfoInterface {
	return &f.info
}
