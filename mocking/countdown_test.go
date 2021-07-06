package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to go!", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, got %q want 4", spySleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

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

		t.Log(spySleeperPrinter.Calls)
		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})

}
