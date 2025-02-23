package main

import (
	"fmt"
	"sort"
	"testing"
)

type SumTest struct {
	testValues []int
	expectedResult int
}

func TestSum(t *testing.T) {
	sumTests := []SumTest {
		{testValues: []int{10, 20, 30}, expectedResult: 10},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
	}

	for index, sumTest := range sumTests {
		t.Run(fmt.Sprintf("Sum #%d", index), func(subT *testing.T) {
			_, total := sortAndTotal(sumTest.testValues)

			if t.Failed() {
				subT.SkipNow()
			}

			if total != sumTest.expectedResult {
				subT.Fatalf("Expected %d, Got %d\n", sumTest.expectedResult, total)
			}
		})

	}
}

func TestSort(t *testing.T) {
	testValues := []int{9,3,7, 2, 9, 1, 4, 6}
	sorted, _ := sortAndTotal(testValues)

	if !sort.IntsAreSorted(sorted) {
		t.Errorf("List is not sorted")
	}
}

func TestSort2(t *testing.T) {
	slices := [][]int {
		{1, 2, 23, 14, 5, 6},
		{4, 5, 6, 8, 10, 11},
		{7, 18, 9, 3, 12, 1},
		{10, 11, 12, 13, 14, 15},
		{1, 2, 3, 4, 5, 6},
	}

	for index, slice := range slices {
		t.Run(fmt.Sprintf("Sort #%d", index), func(subT *testing.T) {
			if index == 1 {
				subT.Skip("Skipping test")
			}

			sorted, _ := sortAndTotal(slice)

			if index == 2 {
				subT.FailNow()
			}

			if !sort.IntsAreSorted(sorted) {
				subT.Errorf("Unsorted slice: %d\n", slice)
			}
		})
	}
}
