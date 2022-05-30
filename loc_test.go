package gofile

import (
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"testing"
)

func TestLocation(t *testing.T) {
	// let's open the first location from where we want to copy some files

	loc1 := &filesystem.Location{
		DirPath: "./loc1/",
	}
	loc2 := &filesystem.Location{
		DirPath: "./loc2/",
	}

	oploc1, _err := OpenLocation(loc1)
	if _err != nil {
		t.Error(_err)
		return
	}

	// now let's identify some files from the loc 1

	file, _err := loc1.FindFile("vasea.txt")
	if _err != nil {
		// failed to find file
		t.Error(_err)
		return
	}

	file, _err := loc2.CopyFile(file, driver.FileDestination{
		Path:     "vasea.txt",
		FileMode: 0751,
	})
	if _err != nil {
		// failed to copy file
		t.Error(_err)
		return
	}

}
