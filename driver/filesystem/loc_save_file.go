package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
	"os"
)

func (l *Location) SaveFile(
	sourceFile driver.FileInterface,
	fileDestPath string,
	perm os.FileMode,
) (bool, error) {
	// Read the file
	var data []byte
	_, _err := sourceFile.Read(data)
	if _err != nil {
		// failed to read...
		return false, _err
	}

	_err = os.WriteFile(fileDestPath, data, perm)
	if _err != nil {
		// failed to write...
		return false, _err
	}
	return true, nil
}
