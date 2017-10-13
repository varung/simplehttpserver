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

  port := ":8000"
  if len(os.Args) > 1 {
    port = ":"+os.Args[1]
  }
  log.Fatal(http.ListenAndServe(port, nil))
}
