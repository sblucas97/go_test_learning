package dependencyInjection

import (
	"bytes"
	"testing"
)

const name string = "Lucas"

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, " + name

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
