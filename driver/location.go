package driver

type LocationInterface interface {
	// Open the location
	Open() (bool, error)

	//Save()
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
