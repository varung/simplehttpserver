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
    // TODO: log like python exactly?
    // 127.0.0.1 - - [13/Oct/2017 02:18:04] "GET /main.go HTTP/1.1" 200 -
    log.Println(r.Method, r.URL.Path)
		w.Header().Add("cache-control", "no-cache")
		fileserver.ServeHTTP(w, r)
	})

  port := "8000"
  if len(os.Args) > 1 {
    port = os.Args[1]
  }
  log.Println("Serving HTTP on 0.0.0.0 port " + port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
