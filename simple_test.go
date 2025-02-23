package main

import (
	"sort"
	"testing"
)

func TestSum(t *testing.T) {
	testValues := []int {10, 20, 30}
	_, sum := sortAndTotal(testValues)

	expected := 60

	if sum != expected {
		t.Errorf("Expected %d, Got %d\n", expected, sum)
	}
}

func TestSort(t *testing.T) {
	testValues := []int{9,3,7, 2, 9, 1, 4, 6}
	sorted, _ := sortAndTotal(testValues)

	if !sort.IntsAreSorted(sorted) {
		t.Errorf("List is not sorted")
	}
}
