package helper

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ReadFile interface {
	Read(b []byte) (int, error)
}

func GetFileContentTypeByPath(filePath string) (string, error) {
	file, _err := os.Open(filePath)
	if _err != nil {
		return "", _err
	}
	defer file.Close()
	return GetFileContentType(file)
}

func GetFileContentType(f ReadFile) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

type FileMetaData struct {
	ContentType string
	Size        int64
	Name        string
	BaseName    string
	Extension   string
}

type FileNameDetails struct {
	Name      string
	BaseName  string
	Extension string
}

func ExtractFileNameDetails(baseName string) FileNameDetails {
	// Try getting the file name and extension

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

	return FileNameDetails{
		Name:      fileName,
		BaseName:  baseName,
		Extension: extension,
	}
}

// Root -> It should return the current process dir path
func Root() (dir string) {
	//currentDirPath, err := os.Getwd()
	path, err := os.Executable()
	if err != nil || path == "" {
		log.Println("Error getting Root Path", err)
		return ""
	}
	currentDirPath := filepath.Dir(path)
	currentDirPath += filepath.FromSlash("/")
	return currentDirPath
}
