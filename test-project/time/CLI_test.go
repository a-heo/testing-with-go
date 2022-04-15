package poker_test

import (
	"strings"
	"testing"
	"time"
	"fmt"
	"github.com/a-heo/testing-with-go/tree/main/test-project/time"
)

type scheduledAlert struct {
	at time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}
type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

var dummySpyAlerter = &SpyBlindAlerter{}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func TestCLI(t *testing.T) {
	
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert {
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}
				got := blindAlerter.alerts[i]

				// amountGot := alert.amount
				// if amountGot != want.expectedAmount {
				// 	t.Errorf("got amount %d, want %d", amountGot, want.expectedAmount)
				// }

				// gotScheduledTime := alert.scheduledAt
				// if gotScheduledTime != want.expectedScheduleTime {
				// 	t.Errorf("got scheduled time %v, want %v", gotScheduledTime, c.expectedScheduleTime)
				// }
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("record chris wins from user input", func(t *testing.T) {

		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()
		
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {

		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})


}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}