package filesystem

import (
	"errors"
	"os"
)

func (l *Location) FileExists(filePath string) (bool, error) {
	_, _err := os.Stat(filePath)
	if _err == nil {
		return true, nil
	}
	if errors.Is(_err, os.ErrNotExist) {
		return false, nil
	}
	return false, _err
}
