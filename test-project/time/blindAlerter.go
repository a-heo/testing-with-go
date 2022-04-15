package poker

import (
	"fmt"
	"os"
	"time"
)

//schedules alerts for blind amounts 
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

//implements blindalerter interface with a function
type BlindAlerterFunc func(duration time.Duration, amount int)

//blindalerterfunc implementation of blindalerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

//schedule alerts and print them to os.Stdout
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}