package driver

import (
	"github.com/google/uuid"
	"os"
	"time"
)

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

//type FileStatusChange func(status *FileStatus)

type FileStatus struct {
	//DirPath  string
	//FilePath string
	//Name     string
	//FileMode os.FileMode
	// in case if it's copied to DB
	//ID *uuid.UUID

	// File ->  Returns back the file itself
	File FileInterface
}

type FileInterface interface {
	//
	Read([]byte) (int, error)
	Status() FileStatus

	Name() string
	Size() int64
	Extension() string
	Path() string
	DirPath() string
	Mode() os.FileMode
	ID() *uuid.UUID

	// TODO: to a function which will read in chunks, not entirely!
	// 		 we need the transfer between the source and destination to be
	//		 - smart
	//		 - low memory
	//		 - fast
}
