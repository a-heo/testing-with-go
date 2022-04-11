package main

import (
	"github.com/a-heo/testing-with-go/tree/main/test-project/commandline"
	"log"
	"net/http"
	"os"
)

//file for db
const dbFileName = "../../game.db.json"

func main() {
	//create file for db, second arg defines permission (read and write or create if it doesn't exist), 3rd arg sets permission for files (all users can read and write) 
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("could not create file system player store, %v", err)
	}
	
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5002", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)		
	}

}