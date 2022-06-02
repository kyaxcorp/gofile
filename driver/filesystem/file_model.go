package filesystem

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	//FPath string // TODO: this var should be removed, because all the necessary information about the file should
	// be stored in the info

	// this is files location
	//location driver.LocationInterface // TODO: ce ce aici nu e native location?
	location *Location // TODO: ce ce aici nu e native location?
	//
	info FileInfo
}
