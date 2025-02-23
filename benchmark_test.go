package main

import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)

func BenchmarkSortBench(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 250

	data := make([]int, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rng.Int()
		}

		b.StartTimer()
		sortAndTotal(data)
	}
}

func BenchmarkSortSubBenchs(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	sizes := []int{10, 100, 250}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array size %v", size), func(subB *testing.B) {
			data := make([]int, size)

			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rng.Int()
				}

				subB.StartTimer()
				sortAndTotal(data)
			}
		})
	}
}
