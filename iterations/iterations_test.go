package iterations

import (
	"strings"
	"testing"
)

const letter string = "a"
const letterLength = 5

func TestIterations (t *testing.T) {
	got := Repeate(letter, letterLength)
	want := "aaaaa"

	if strings.Compare(got, want) != 0 {
		t.Errorf("Expects aaaaa, but got %q", got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate(letter, letterLength)
	}
}