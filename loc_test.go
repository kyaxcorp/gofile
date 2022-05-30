package gofile

import (
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"testing"
)

func TestLocation(t *testing.T) {
	loc1, _err := OpenLocation(&filesystem.Location{
		DirPath: "./",
	})
	if _err != nil {
		t.Error(_err)
		return
	}

}
