package main

import (
	"fmt"
	"log"
	"os"
	poker 	"github.com/a-heo/testing-with-go/tree/main/test-project/time"
)

const dbFileName = "../../game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	
	if err!= nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	fmt.Println("let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	// game := poker.NewCLI(store, os.Stdin)
	// game.PlayPoker()
	//starts the game w/ alert printed 
	// poker.NewCLI(store, os.Stdin, os.Stdout, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()

	cli.PlayPoker()

}