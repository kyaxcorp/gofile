package filesystem

import "github.com/kyaxcorp/gofile/driver"

func (l *Location) MoveFile(
	file driver.FileInterface,
	dest driver.FileDestination,
) (*driver.FileStatus, error) {
	return nil, nil
}
