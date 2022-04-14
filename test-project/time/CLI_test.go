package poker_test

import (
	"github.com/a-heo/testing-with-go/tree/main/test-project/commandline"
	"testing"
	"strings"
)

func TestCLI(t *testing.T) {
	
	t.Run("record chris wins from user input", func(t *testing.T) {

		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {

		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})


}