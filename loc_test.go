package gofile

import (
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"testing"
)

func TestLocation(t *testing.T) {
	// let's open the first location from where we want to copy some files
	loc1, _err := OpenLocation(&filesystem.Location{
		DirPath: "./",
	})
	if _err != nil {
		t.Error(_err)
		return
	}

	// now let's identify some files from the loc 1

}
