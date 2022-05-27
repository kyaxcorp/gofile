package filesystem

type Location struct {
	// TODO: we should add some properties here?! Disk, partition at least?!
	// don't think so, because it's the OS filesystem
	// or it would be very good to have a specific PATH as prefix!!! for not saving the entire path of the file!

	DirPath string
}
