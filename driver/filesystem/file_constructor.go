package filesystem

func newFile(filePath string) (*File, error) {
	// Generate file info
	fInfo, _err := newFileInfo(filePath)
	if _err != nil {
		return nil, _err
	}

	// generate the file itself
	f := &File{
		FPath: filePath,
		info:  *fInfo,
	}

	// return
	return f, nil
}
