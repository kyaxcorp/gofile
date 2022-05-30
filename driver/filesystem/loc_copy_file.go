package filesystem

import (
	"github.com/kyaxcorp/gofile"
	"github.com/kyaxcorp/gofile/err"
	"os"
)

func (l *Location) CopyFile(
	file gofile.FileInterface,
	dest gofile.FileDestination,
) (gofile.FileInterface, error) {
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

	return newFile(dest.Path)
}