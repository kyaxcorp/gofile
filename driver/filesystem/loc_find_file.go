package filesystem

import (
	"github.com/kyaxcorp/gofile"
	"github.com/kyaxcorp/gofile/err"
)

func (l *Location) FindFile(filePath string) (gofile.FileInterface, error) {
	// Check if file exists
	// if yes then return the file

	// Reads the whole file...
	// but we want to retrieve a just an identification of the file
	// and when reading will be necessary, it will be done!

	//https://golang.cafe/blog/golang-read-file-example.html

	// File the file!
	// Check if exists

	fullFilePath := l.GetFilePath(filePath)

	exists, _err := l.FileExists(filePath)
	if _err != nil {
		return nil, _err
	}
	if !exists {
		return nil, err.ErrFileDoesntExist
	}

	// Return the current file
	return newFile(fullFilePath)
}
