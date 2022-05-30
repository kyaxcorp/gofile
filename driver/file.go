package driver

import (
	"github.com/google/uuid"
	"os"
	"time"
)

type FileInterface interface {
	//
	Read() ([]byte, error)
	Info() FileInfoInterface // TODO: we can return all the functions from the current file over there

	// TODO: we need a function which will return the structure

	// Get location of the file
	Location() LocationInterface
}

type Any = interface{}

type FileInfoInterface interface {
	// ToStruct ->  Converts to a struct
	ToStruct() FileInfo

	// ID - is available only when saving to a storage which supports File Identification
	ID() *uuid.UUID
	//
	Name() string
	FullName() string
	//
	Path() string
	DirPath() string

	ContentType() string
	Size() int64
	Extension() string

	Mode() os.FileMode
	//
	UpdatedAt() time.Time
	CreatedAt() time.Time
	// Sys - other system info for this file
	Sys() Any
	// TODO: add function Sys
	// 		which will return file owner, permissions for other OS'ses and other storage platforms?!
}

type FileInfo struct {
	ID          *uuid.UUID
	Name        string
	FullName    string
	Path        string
	DirPath     string
	ContentType string
	Size        int64
	Extension   string
	Mode        os.FileMode
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Sys         Any
}

// FileDestination -> is used when copying or moving to a new destination
type FileDestination struct {
	Path     string
	FileMode os.FileMode

	// if you want to receive in real time the status (how much left to copy)
	// put a callback handler over here
	OnCopy FileOnCopy

	// If copied to DB
	// nothing is needed!
}

type FileOnCopy func(status FileCopyStatus)

type FileCopyStatus struct {
	PercentageCopied int8
	BytesCopied      int64
	LeftBytes        int64
	// when it will finish
	FinishETA   time.Time
	LeftSeconds int64
}

//type FileStatusChange func(status *FileInfo)
