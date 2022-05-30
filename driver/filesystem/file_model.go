package filesystem

import "github.com/kyaxcorp/gofile"

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	FPath string

	// this is files location
	location gofile.LocationInterface
	//
	info FileInfo
}
