package main

import (
	"fmt"
	"sort"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, v := range sorted {
		total += v
		//total++
	}
	return
}

func main() {
	num := []int {8, 2, 5, 3, 6, 1, 4, 7}
	sorted, total := sortAndTotal(num)
	fmt.Printf("Sorted: %v\n", sorted)
	fmt.Printf("Total: %d\n", total)
}
