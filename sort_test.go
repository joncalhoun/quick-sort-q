package sort

import (
	"math/rand"
	"testing"
)

func benchmarkSortAlgorithm(b *testing.B, listSize int, sortFn func([]int) []int) {
	list := make([]int, listSize)
	var prng = rand.New(rand.NewSource(42))
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()
		sortFn(list)
	}
}

func testSortAlgorithm(t *testing.T, listSize int, sortFn func([]int) []int) {
	list := make([]int, listSize)
	var prng = rand.New(rand.NewSource(42))
	for i := range list {
		list[i] = prng.Int()
	}
	list = sortFn(list)
	last := list[0]
	for _, v := range list {
		if v < last {
			t.Errorf("list is not sorted")
		}
	}
}

func TestQuickSort_10_000_000(t *testing.T) {
	testSortAlgorithm(t, 10000000, quickSort)
}
func TestQuickSort_1_000_000(t *testing.T) {
	testSortAlgorithm(t, 1000000, quickSort)
}
func TestQuickSort_10_000(t *testing.T) {
	testSortAlgorithm(t, 10000, quickSort)
}

func TestQuickSortConcurrent_10_000_000(t *testing.T) {
	testSortAlgorithm(t, 10000000, quickSortConcurrent)
}
func TestQuickSortConcurrent_1_000_000(t *testing.T) {
	testSortAlgorithm(t, 1000000, quickSortConcurrent)
}
func TestQuickSortConcurrent_10_000(t *testing.T) {
	testSortAlgorithm(t, 10000, quickSortConcurrent)
}

func TestSplitSort_10_000_000(t *testing.T) {
	testSortAlgorithm(t, 10000000, splitSort)
}
func TestSplitSort_1_000_000(t *testing.T) {
	testSortAlgorithm(t, 1000000, splitSort)
}
func TestSplitSort_10_000(t *testing.T) {
	testSortAlgorithm(t, 10000, splitSort)
}

func BenchmarkQuickSort_10_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000000, quickSort)
}
func BenchmarkQuickSort_1_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 1000000, quickSort)
}
func BenchmarkQuickSort_10_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000, quickSort)
}

func BenchmarkQuickSortConcurrent_10_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000000, quickSortConcurrent)
}
func BenchmarkQuickSortConcurrent_1_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 1000000, quickSortConcurrent)
}
func BenchmarkQuickSortConcurrent_10_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000, quickSortConcurrent)
}

func BenchmarkSplitSort_10_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000000, splitSort)
}
func BenchmarkSplitSort_1_000_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 1000000, splitSort)
}
func BenchmarkSplitSort_10_000(b *testing.B) {
	benchmarkSortAlgorithm(b, 10000, splitSort)
}
