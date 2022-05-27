package filesystem

func (l *Location) Open() (bool, error) {
	// if it's the current file system... then we should return ok...
	// but still, it cannot be ok if it's a separate disk or other partition
	// TODO: we should analyze here...
	return true, nil
}
