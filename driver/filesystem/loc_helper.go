package filesystem

import (
	"path/filepath"
)

func (l *Location) GetFilePath(filePath string) string {
	return l.DirPath + filepath.FromSlash("/") + filePath
}
