package filesystem

type File struct {
	// Where it's physically located
	// should we store full paths, or simply part of the path
	FPath string

	info FileInfo
}
