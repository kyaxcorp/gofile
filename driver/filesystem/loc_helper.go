package filesystem

import (
	"path/filepath"
)

func (l *Location) GetFilePath(filePath string) string {
	p := ""

	// Get current app path!
	// if a full path is been provided , then the current app path is not considered!
	// we should not take in consideration the CWD current working directory!

	//helper.Root()

	// if writing "./storage" -> it should take the current working directory
	// if writing "storage" -> it should take the current working directory
	// if writing "/storage/" -> it should take the full path starting from root!
	// if

	if l.DirPath != "" {
		p = p + l.DirPath + filepath.FromSlash("/")
	}
	p = p + filePath
	return p
}
