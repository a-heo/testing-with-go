package poker

import (
	"bufio"
	"strings"
	"io"
)

//helps players through game of poker
type CLI struct {
	playerStore PlayerStore
	in 	*bufio.Scanner
}

//constructor function for taking in inputs  
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in: bufio.NewScanner(in),
	}
}

//starts game
func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}