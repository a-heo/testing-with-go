package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
//finish up mocking 