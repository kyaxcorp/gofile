package tests

import (
	"log"
	"strings"
	"testing"
)

func TestExtractName(t *testing.T) {
	baseName := "helloworld.txt"
	extensionSep := "."
	var fileName, extension string
	if strings.Contains(baseName, extensionSep) {
		fileBaseNameSplit := strings.Split(baseName, extensionSep)
		// Recreate the file name without extension
		// Take from 0 to the pre last element
		fileName = strings.Join(fileBaseNameSplit[:len(fileBaseNameSplit)-1], extensionSep)
		// take the last one!
		extension = fileBaseNameSplit[len(fileBaseNameSplit)-1]
	} else {
		// No extension available
		fileName = baseName
	}
	log.Println(fileName, extension)
}
