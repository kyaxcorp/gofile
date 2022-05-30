package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/err"
	"os"
)

func (l *Location) CopyFile(
	file driver.FileInterface,
	dest driver.FileDestination,
) (driver.FileInterface, error) {
	// Read the file
	data, _err := file.Read()
	if _err != nil {
		// failed to read...
		return nil, _err
	}

	if dest.Path == "" {
		return nil, err.ErrDestinationPathIsEmpty
	}

	destPath := l.GetFilePath(dest.Path)

	fileMode := file.Info().Mode()

	// override the current files one!
	if dest.FileMode != 0 {
		fileMode = dest.FileMode
	}

	// if no file mode has been set or retrieved
	if fileMode == 0 {
		fileMode = 0751
	}

	_err = os.WriteFile(destPath, data, fileMode)
	if _err != nil {
		// failed to write...
		return nil, _err
	}

	// we should give the path for that location, nothing more! (the one that's indicated by the user)
	return l.newFile(dest.Path)
}
