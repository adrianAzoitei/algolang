package mergesort_test

import (
	"algorithms/mergesort"
	"math/rand/v2"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// arrange
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	shuffled := make([]int, len(want))
	perm := rand.Perm(len(want))
	for i, v := range perm {
		shuffled[v] = want[i]
	}
	t.Log(shuffled)

	if reflect.DeepEqual(want, shuffled) {
		t.Error("Shuffling gone wrong")
	}

	// act
	got, err := mergesort.MergeSort(shuffled)
	// assert
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
	t.Log(got)
}
