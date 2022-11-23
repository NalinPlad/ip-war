package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"net"
	"io"

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

	if lbname == "" {
		io.WriteString(w, "please supply a name\n");
		return
	}

	// Open DB
	db, err := sql.Open("sqlite3", "./IP_Leaderboard.db")
	
	if err != nil {
			log.Fatal(err)
	}
  
	defer db.Close()

	// SQL Query to INSERT if data does not exist of REPLACE if it already exists
	var sts = fmt.Sprintf("INSERT OR REPLACE INTO address(IP, lbname) VALUES('%s','%s');", ip, lbname)

	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(fmt.Sprintf("registered %s for %s", ip, lbname))
}

func GetData (w http.ResponseWriter, r *http.Request) {	
	// Open DB
	db, err := sql.Open("sqlite3", "./IP_Leaderboard.db")
	
	if err != nil {
		log.Fatal(err)
	}
  
	defer db.Close()

	// Get data from db
	rows, err := db.Query("SELECT * FROM address ORDER BY IP DESC LIMIT 100")

	if err != nil {
		log.Fatal(err)
	}

	var IP string
	var lbname string

	// Array of 100 arrays of 2 strings
	var IP_Array [100][2]string

	var i = 0

	for rows.Next() {
		err = rows.Scan(&IP, &lbname)

		if err != nil {
			log.Fatal(err)
		}
		
		IP_Array[i] = [2]string{IP, lbname}
		i++

		fmt.Println(fmt.Sprintf("%s owns %s", lbname, IP))
	}

	// Convert Arr to JSON
	fmt.Println(IP_Array)

	var IP_JSON, _ = json.Marshal(IP_Array)
	
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, string(IP_JSON))
}

func main() {

	// Main Page
	http.Handle("/", http.FileServer(http.Dir("./public")))

  // Register IP
	http.HandleFunc("/play", Handle)

	// Get data
	http.HandleFunc("/data", GetData)

	// Serve http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
