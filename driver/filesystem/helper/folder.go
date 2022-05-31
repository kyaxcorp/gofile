package helper

import "os"

func FolderExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func MkDir(path string, perm ...os.FileMode) error {
	var p os.FileMode
	p = 0751
	if len(perm) > 0 {
		p = perm[0]
	}
	return os.MkdirAll(path, p)
}
