package poker

import (
	"io"
	"bufio"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in 	io.Reader
}

func (cli *CLI) PlayPoker() {
	//read input from io reader
	reader := bufio.NewScanner(cli.in)
	//read up to a newline
	reader.Scan()
	//text method returns string the scanner read to
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}