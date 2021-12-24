package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go"
)

type ConfigurableSleeper struct {
	delay time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.delay)
}

type Sleeper interface {
	Sleep()
}

func CountDown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}
