package filesystem

import (
	"path/filepath"
)

func (l *Location) GetFilePath(filePath string) string {
	p := ""
	if l.DirPath != "" {
		p = p + l.DirPath + filepath.FromSlash("/")
	}
	p = p + filePath
	return p
}
