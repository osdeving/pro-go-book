package main

import (
	"log"
	"sort"
)

func sortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), " ▶️ 🔧 sortAndTotal 🔧 : ", log.Flags() | log.Lmsgprefix)

	logger.Printf("Invoked with %v values\n", len(vals))

	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)

	logger.Printf("Sorted data: %v\n", sorted)

	for _, v := range sorted {
		total += v
		//total++
	}

	logger.Printf("Total: %d\n", total)

	return
}

func main() {
	num := []int{8, 2, 5, 3, 6, 1, 4, 7}
	sorted, total := sortAndTotal(num)
	log.Printf("Sorted: %v\n", sorted)
	log.Printf("Total: %d\n", total)
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.LstdFlags)
}
