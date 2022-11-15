package main

import (
	"log"
	"net/http"
	"fmt"
)


func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("registering %s", r.RemoteAddr))
}

func main() {
	// Main Page
	http.HandleFunc("/play", Handle)
	http.Handle("/", http.FileServer(http.Dir("./public")))
  
	// Serve http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
