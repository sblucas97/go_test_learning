package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		assertCorrectMessage(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got []int, want []int, numbers ...[]int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 2 arrays of 5 numbers", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumAll(numbers, numbers)
		want := []int {15, 15}

		assertCorrectMessage(t, got, want, numbers, numbers)
	})

	t.Run("collection of any arrays numbers of 5 numbers", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumAll(numbers, numbers, numbers, numbers, numbers, numbers)
		want := []int {15, 15, 15, 15, 15, 15}

		assertCorrectMessage(t, got, want, numbers, numbers, numbers, numbers, numbers, numbers)
	})
}

func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got []int, want []int, numbers ...[]int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("a array of 5 numbers", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := SumAllTails(numbers)
		want := []int {20}

		assertCorrectMessage(t, got, want, numbers)
	})

	t.Run("collection of any arrays numbers of 5 numbers", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumAllTails(numbers, numbers, numbers, numbers, numbers, numbers)
		want := []int {14, 14, 14, 14, 14, 14}

		assertCorrectMessage(t, got, want, numbers, numbers, numbers, numbers, numbers, numbers)
	})
	
	t.Run("collection of arrays with a empty one", func(t* testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		emptyArray := []int{}

		got := SumAllTails(numbers, emptyArray)
		want := []int {14, 0}

		assertCorrectMessage(t, got, want, numbers, emptyArray)
	})
}