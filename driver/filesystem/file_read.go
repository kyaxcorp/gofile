package filesystem

import (
	"os"
)

func (f *File) Read() ([]byte, error) {
	//return os.ReadFile(f.FPath)
	physicalPath := f.location.GetFilePath(f.Info().Path())
	return os.ReadFile(physicalPath)
}
