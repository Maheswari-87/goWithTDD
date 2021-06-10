package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Testing 3,2,1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		//spySleeper := &SpySleeper{}
		Countdown(buffer, &CountDownOperationsSpy{})

		got := buffer.String()

		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
	t.Run("Sleep before every print,", func(t *testing.T) {
		spySleepPrinter := &CountDownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got calls %v", want, spySleepPrinter.Calls)
		}

	})
}

type CountDownOperationsSpy struct {
	Calls []string
}

func (c *CountDownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}
func (c *CountDownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	sleptDuration time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.sleptDuration = duration
}
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.sleptDuration != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.sleptDuration)
	}

}
