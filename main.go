package main

import (
	"net/http"
	"sync"
	"log"
	"os"
)

var mut = sync.Mutex{}

func main() {

	fileserver := http.FileServer(http.Dir("./"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("cache-control", "no-cache")
		fileserver.ServeHTTP(w, r)
	})
  log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
