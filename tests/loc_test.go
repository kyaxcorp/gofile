package tests

import (
	"github.com/kyaxcorp/gofile"
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	// let's open the first location from where we want to copy some files

	loc1 := &filesystem.Location{
		DirPath: "./",
		//DirPath: ".",
	}

	loc2 := &filesystem.Location{
		DirPath: "./",
		//DirPath: ".",
	}

	// now let's identify some files from the loc 1

	file, _err := loc1.FindFile("testfile.md")
	if _err != nil {
		// failed to find file
		t.Error(_err)
		return
	}

	newFile, _err := loc2.CopyFile(file, gofile.FileDestination{
		Path:     "testfile2.md",
		FileMode: 0751,
	})
	if _err != nil {
		// failed to copy file
		t.Error(_err)
		return
	}

	log.Println(newFile.Info().ToStruct())
}
