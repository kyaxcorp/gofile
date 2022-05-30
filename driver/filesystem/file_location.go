package filesystem

import "github.com/kyaxcorp/gofile"

func (f *File) Location() gofile.LocationInterface {
	return f.location
}
