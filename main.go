package main

import (
	"log"
	"net/http"
)


func main() {
	// Main Page
	http.Handle("/", http.FileServer(http.Dir("./public")))
  
	// Serve http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
