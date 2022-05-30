package filesystem

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/driver/filesystem/helper"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	path        string
	fullName    string
	name        string
	contentType string
	size        int64
	extension   string
	dirPath     string
	mode        os.FileMode
	createdAt   time.Time
	updatedAt   time.Time
	//
	sys driver.Any
}

func NewFileInfo(filePath string) (*FileInfo, error) {
	fileInfo, _err := os.Stat(filePath)
	if _err != nil {
		return nil, _err
	}

	baseName := fileInfo.Name()
	// Get naming details
	fileNameDetails := helper.ExtractFileNameDetails(baseName)

	// Get content type
	contentType, _err := helper.GetFileContentTypeByPath(filePath)
	if _err != nil {
		return nil, _err
	}

	fInfo := &FileInfo{
		path:        filePath,
		name:        fileNameDetails.Name,
		fullName:    fileNameDetails.BaseName,
		size:        fileInfo.Size(),
		contentType: contentType,
		extension:   fileNameDetails.Extension,
		updatedAt:   fileInfo.ModTime(),
		//createdAt:   time.Now(), // TODO:
		dirPath: filepath.Dir(filePath),
		mode:    fileInfo.Mode(),
		sys:     fileInfo.Sys(),
	}
	return fInfo, nil
}

func (f *FileInfo) ToStruct() driver.FileInfo {
	return driver.FileInfo{
		ID:          f.ID(),
		Name:        f.Name(),
		FullName:    f.FullName(),
		Path:        f.Path(),
		DirPath:     f.DirPath(),
		ContentType: f.ContentType(),
		Size:        f.Size(),
		Extension:   f.Extension(),
		Mode:        f.Mode(),
		UpdatedAt:   f.UpdatedAt(),
		CreatedAt:   f.CreatedAt(),
		Sys:         f.Sys(),
	}
}

func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) Sys() driver.Any {
	return f.sys
}

func (f *FileInfo) FullName() string {
	return f.fullName
}

func (f *FileInfo) Size() int64 {
	return f.size
}

func (f *FileInfo) Extension() string {
	return f.extension
}

func (f *FileInfo) Path() string {
	return f.path
}

func (f *FileInfo) DirPath() string {
	return f.dirPath
}

func (f *FileInfo) UpdatedAt() time.Time {
	return f.updatedAt
}

func (f *FileInfo) CreatedAt() time.Time {
	return f.createdAt
}

func (f *FileInfo) ContentType() string {
	return f.contentType
}

func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *FileInfo) ID() *uuid.UUID {
	return nil
}
