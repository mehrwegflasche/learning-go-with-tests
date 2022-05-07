package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSumArrays(t *testing.T) {
	assertIfMatches := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d , expected %d", got, want)
		}
	}

	t.Run("Check if the sum of 5 matches the passed total", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertIfMatches(t, got, want)
	})

	t.Run("Check if the sum of 3 matches the passed total", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertIfMatches(t, got, want)
	})

}

func TestSumAllArrays(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted%v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
}
