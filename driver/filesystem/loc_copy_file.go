package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/err"
	"os"
)

func (l *Location) CopyFile(
	file driver.FileInterface,
	dest driver.FileDestination,
) (*driver.FileInfo, error) {
	// Read the file
	var data []byte
	_, _err := file.Read(data)
	if _err != nil {
		// failed to read...
		return nil, _err
	}

	if dest.Path == "" {
		return nil, err.ErrDestinationPathIsEmpty
	}

	// set default file mode in case if it's not set
	if dest.FileMode == 0 {
		dest.FileMode = 0751
	}

	_err = os.WriteFile(dest.Path, data, dest.FileMode)
	if _err != nil {
		// failed to write...
		return nil, _err
	}

	// TODO: generate the new file from the new location

	return &driver.FileInfo{
		DirPath: dest.Path,
	}, nil
}
