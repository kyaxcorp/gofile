package gofile

import (
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/err"
)

func OpenLocation(driver driver.LocationInterface) (*Location, error) {
	/*
		1. connect to this location!?
		2. check if it's ok!
		3. return back the instance!
	*/

	ok, _err := driver.Open()
	if _err != nil {
		return nil, _err
	}
	if !ok {
		return nil, err.ErrFailedToOpenLocation
	}

	return &Location{
		driver: driver,
	}, nil
}
