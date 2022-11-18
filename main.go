package main

import (
	"log"
	"net/http"
	"fmt"
	"net"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func Handle(w http.ResponseWriter, r *http.Request) {
	// Get IP
	var ip,_,err = net.SplitHostPort(r.RemoteAddr)
	
	if err != nil {
			log.Fatal(err)
	}

	// Get lbname
	var lbname = r.URL.Query().Get("name")

	// Open DB
	db, err := sql.Open("sqlite3", "./IP_Leaderboard.db")
	
	if err != nil {
			log.Fatal(err)
	}
  
	defer db.Close()

	var sts = fmt.Sprintf("INSERT OVERWRITE TABLE address(IP, lbname) VALUES('%s','%s');", ip, lbname)	

	_, err = db.Exec(sts)

	if err != nil {
			log.Fatal(err)
	}
	
	log.Println(fmt.Sprintf("registered %s for %s", ip, lbname))
}

func main() {

	// Main Page
	http.HandleFunc("/play", Handle)
	http.Handle("/", http.FileServer(http.Dir("./public")))
  
	// Serve http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
