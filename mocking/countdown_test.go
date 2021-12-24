package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// can name return var to avoid having to initialize them
// implements io.Writer
func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("should print out correct sequence", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperation{}

		CountDown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should alternate between write & sleep", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperation{}

		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", spyTime.durationSlept, sleepTime)
	}
}
