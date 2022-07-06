package main

import (
	"log"
	"path/filepath"
)

func main() {

	path1 := "./"
	path2 := "vasea"

	log.Println(filepath.Base(path1))
	log.Println(filepath.Dir(path1))
	log.Println(filepath.Clean(path1))
	log.Println(filepath.Abs(path1))
	log.Println(filepath.Ext(path1))

	log.Println("path 2")
	log.Println(filepath.Base(path2))
	log.Println(filepath.Dir(path2))
	log.Println(filepath.Clean(path2))
	log.Println(filepath.Abs(path2))
	log.Println(filepath.Ext(path2))

}
