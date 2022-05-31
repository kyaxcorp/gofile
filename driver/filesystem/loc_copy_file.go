package filesystem

import (
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/driver/filesystem/helper"
	"github.com/kyaxcorp/gofile/err"
	"os"
	"path/filepath"
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

	// let's generate files path
	destPath := ""
	if dest.FilePath != "" {
		// Check if file path is indicated! it overrides everything
		destPath = dest.FilePath
	} else if dest.DirPath != "" {
		// check if DirPath is indicated, it overrides the default one
		// take the file name
		destPath = dest.DirPath + filepath.FromSlash("/") + file.Info().FullName()
	} else {
		// if none of the optional params are indicated, it will take the current files
		// Base Path and recreate on the new location
		// Get file path
		destPath = file.Info().Path()
	}

	if destPath == "" {
		return nil, err.ErrDestinationPathIsEmpty
	}

	if dest.FileName != "" {
		tmpDir := filepath.Dir(destPath)
		// Override the file name...
		destPath = tmpDir + filepath.FromSlash("/") + dest.FileName
		//tmpFile := filepath.Base(destPath)
	}

	fileDestPath := l.GetFilePath(destPath)

	fileMode := file.Info().Mode()

	// override the current files one!
	if dest.FileMode != 0 {
		fileMode = dest.FileMode
	}

	// if no file mode has been set or retrieved
	if fileMode == 0 {
		fileMode = 0751
	}

	// But before creating the file, we should check if the Base Dir exists
	// If it doesn't we should create it
	baseDir := filepath.Dir(fileDestPath)
	if !helper.FolderExists(baseDir) {
		// Create the path recursively
		_err := helper.MkDir(baseDir)
		if _err != nil {
			// failed to create folder...
			return nil, _err
		}
	}

	_err = os.WriteFile(fileDestPath, data, fileMode)
	if _err != nil {
		// failed to write...
		return nil, _err
	}

	// we should give the path for that location, nothing more! (the one that's indicated by the user)
	return l.newFile(destPath)
}
