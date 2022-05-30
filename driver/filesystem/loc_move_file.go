package filesystem

import (
	"github.com/kyaxcorp/gofile"
)

func (l *Location) MoveFile(
	file gofile.FileInterface,
	dest gofile.FileDestination,
) (gofile.FileInterface, error) {
	return nil, nil
}
