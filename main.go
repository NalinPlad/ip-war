package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, "+req.RemoteAddr+"")
		log.Println(req.RemoteAddr)
	}

	http.HandleFunc("/", Handler)
  log.Println("http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
