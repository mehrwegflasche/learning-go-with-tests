package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	FinalGreeting  = "Go!"
	countDownStart = 3
)

func Countdown(out io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()

	}
	fmt.Fprint(out, FinalGreeting)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

type Sleeper interface {
	Sleep()
}

type SpySleeperPrinter struct {
	Calls []string
}

func (s *SpySleeperPrinter) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeperPrinter) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

const write = "write"
const sleep = "sleep"
