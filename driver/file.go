package driver

type FileInterface interface {
	//
	Read([]byte) (int, error)
	// TODO: to a function which will read in chunks, not entirely!
	// 		 we need the transfer between the source and destination to be
	//		 - smart
	//		 - low memory
	//		 - fast
}
