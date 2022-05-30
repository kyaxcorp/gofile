package gofile

/*
	Names:
	1. Storage
	2. Location
	3. Filesystem
	4.
*/

type LocationInterface interface {
	// Open the location
	Open() (bool, error)

	CopyFile(file FileInterface, dest FileDestination) (FileInterface, error)
	MoveFile(file FileInterface, dest FileDestination) (FileInterface, error)

	// DeleteFile - File status will be used for deletion
	DeleteFile(file FileInterface) error

	//Save() // should not be used...
	//Copy()
	//Move()
	//Delete()
	//Create()
	//Touch()
	//List()

	// Do we need directory manager?!
	// if it's called file manager, then it should be related only to files...
	// how they are saved, should not interest the user?!
	// well, sometimes the user will want to indicate where to save a specific file...
	// but that's other thing, it's not related to making a directory or deleting it...
	// the user should be interested in saving the files and getting them back! that's it!
	//MkDir()
	//DeleteDir()
}
