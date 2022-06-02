package filesystem

import "github.com/kyaxcorp/gofile/driver"

// FileFromInfo -> it recreates the file struct/object from driver.FileInfo so further manipulations can be made
func (l *Location) FileFromInfo(fileInfo driver.FileInfo) (driver.FileInterface, error) {
	nativeFileInfo := DriverFileInfoToNative(fileInfo)

	f := &File{
		info:     nativeFileInfo,
		location: l,
	}

	// return
	return f, nil
}

// TODO:
// FileFromInfoMap -> it recreates the file struct/object from driver.FileInfo so further manipulations can be made
func (l *Location) FileFromInfoMap(fileInfo map[string]interface{}) (driver.FileInterface, error) {
	return nil, nil
}

// TODO:
// FileFromInfoMap -> it recreates the file struct/object from driver.FileInfo so further manipulations can be made
func (l *Location) FileFromInfoJson(fileInfo string) (driver.FileInterface, error) {
	return nil, nil
}
