package main

import (
	"net/http"
	"sync"
	"log"
	"os"
)

var mut = sync.Mutex{}

func main() {

	// TODO: change to flag
	var files_root string = "./"
	fileserver := http.FileServer(http.Dir(files_root))
	log.Fatal(http.ListenAndServe(":"+os.Args[1], fileserver))
}
