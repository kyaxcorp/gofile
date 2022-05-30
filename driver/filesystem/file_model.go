package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
)

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	FPath string

	// this is files location
	location driver.LocationInterface
	//
	info FileInfo
}
