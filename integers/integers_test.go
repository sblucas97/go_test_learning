package integers

import "testing"

func TestIntegers (t *testing.T) {
	got := Add(10, 20)
	want := 30

	if got != want {
		t.Errorf("%v got is different than %v", got, want)
	}
}