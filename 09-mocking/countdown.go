package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {

}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}
