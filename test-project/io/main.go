package main

import (
	"log"
	"net/http"
	"os"
)

//file for db
const dbFileName = "game.db.json"

func main() {
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	//create file for db, second arg defines permission (read and write or create if it doesn't exist), 3rd arg sets permission for files (all users can read and write) 
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	// store := &FileSystemPlayerStore{db}
	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5002", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)		
	}

	// log.Fatal(http.ListenAndServe(":5003", server))
}