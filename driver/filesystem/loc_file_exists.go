package filesystem

import (
	"errors"
	"os"
)

func (l *Location) FileExists(filePath string) (bool, error) {
	fullFilePath := l.GetFilePath(filePath)
	_, _err := os.Stat(fullFilePath)
	if _err == nil {
		// File exists
		return true, nil
	}
	if errors.Is(_err, os.ErrNotExist) {
		return false, os.ErrNotExist
	}
	return false, _err
}
