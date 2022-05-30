package filesystem

import (
	"os"
)

func (f *File) Read() ([]byte, error) {
	return os.ReadFile(f.FPath)
}
