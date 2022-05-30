package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
)

func (f *File) Location() driver.LocationInterface {
	return f.location
}
