package poker

import (
	"bufio"
	"io"
	"strings"
	"fmt"
	"strconv"
)

//construct Game with existing dependencies and interpret user input as method invocations for Game 
type CLI struct {
	in *bufio.Scanner
	out io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in: bufio.NewScanner(in),
		out: out,
		game: game,
	}
}

const PlayerPrompt = "please enter the number of players: "

func (cli *CLI) PlayPoker() {
		fmt.Fprint(cli.out, PlayerPrompt)

		numOfPlayersInput := cli.readLine()
		numOfPlayers, _ := strconv.Atoi(strings.Trim(numOfPlayersInput, "\n"))

		cli.game.Start(numOfPlayers)

		winnerInput := cli.readLine()
		winner := extractWinner(winnerInput)
		cli.game.Finish(winner)
	
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

// //helps players through game of poker
// type CLI struct {
// 	playerStore PlayerStore
// 	in 	*bufio.Scanner
// 	out	io.Writer
// 	alerter BlindAlerter
// }

// //constructor function for taking in inputs when using CLI 
// func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
// 	return &CLI{
// 		playerStore: store,
// 		in: bufio.NewScanner(in),
// 		out: out,
// 		alerter: alerter,
// 	}
// }

// //starts game
// func (cli *CLI) PlayPoker() {
// 	fmt.Fprint(cli.out, PlayerPrompt)

// 	//convert string to int from input
// 	numberOfPlayers, _ := strconv.Atoi(cli.readLine())
// 	//accept a number of players
// 	cli.scheduleBlindAlerts(numberOfPlayers)

// 	userInput := cli.readLine()
// 	cli.playerStore.RecordWin(extractWinner(userInput))
// }

// func (cli *CLI) scheduleBlindAlerts(numOfPlayers int) {
// 	//calculate time to use to add to blindTime when iterating over blind amounts
// 	blindIncrement := time.Duration(5 + numOfPlayers) * time.Minute

// 	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
// 	blindTime := 0 * time.Second
// 	for _, blind := range blinds {
// 		cli.alerter.ScheduleAlertAt(blindTime, blind)
// 		blindTime = blindTime + blindIncrement
// 	}
// }

// func extractWinner(userInput string) string {
// 	return strings.Replace(userInput, " wins", "", 1)
// }

// func (cli *CLI) readLine() string {
// 	cli.in.Scan()
// 	return cli.in.Text()
// }