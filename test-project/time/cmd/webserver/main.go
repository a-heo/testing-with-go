package main

import (
	poker "github.com/a-heo/testing-with-go/tree/main/test-project/commandline"
	"log"
	"net/http"
)

//file for db
const dbFileName = "../../game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5003", server); err != nil {
		log.Fatalf("could not listen on port 5003 %v", err)		
	}

}