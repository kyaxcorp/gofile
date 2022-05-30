package err

import "errors"

var (
	ErrFileDoesntExist        = errors.New("file doesn't exist")
	ErrFailedToOpenLocation   = errors.New("failed to open location")
	ErrDestinationPathIsEmpty = errors.New("destination path is empty")
)
