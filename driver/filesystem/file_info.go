package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
)

func (f *File) Info() driver.FileInfoInterface {
	return &f.info
}
