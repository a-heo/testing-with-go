package main

import (
	"fmt"
	"log"
	"os"
	poker 	"github.com/a-heo/testing-with-go/tree/main/test-project/commandline"
)

const dbFileName = "../../game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	
	if err!= nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()

}