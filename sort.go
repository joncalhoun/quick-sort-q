package sort

import (
	"math/rand"
	"sync"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func quickSortConcurrent(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	var wg sync.WaitGroup
	wg.Add(2)

	// Go down the rabbit hole
	go func() {
		quickSort(a[:left])
		wg.Done()
	}()
	go func() {
		quickSort(a[left+1:])
		wg.Done()
	}()
	wg.Wait()

	return a
}

func splitSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	//used as an effectively random partition value
	mid := a[(len(a)-1)/2]

	//The capacity parameters for these slices could be tuned for different workloads
	//used to hold values > partition value
	b := make([]int, 0, len(a)/4)
	//used to hold values == partition value
	c := make([]int, 0, len(a)/100)

	//number of elements removed from primary slice
	dels := 0
	for i := 0; i < len(a); i++ {
		if a[i] > mid {
			b = append(b, a[i])
			dels++
			continue
		}
		if a[i] < mid {
			if dels > 0 {
				a[i-dels] = a[i]
			}
			continue
		}
		dels++
		c = append(c, a[i])
	}

	//sort the higher partition in a different go-routine
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		splitSort(b)
		wg.Done()
	}()

	//append the partition values to what's left of the primary slice
	//(should be "stable", so equal values don't change order)
	a = append(splitSort(a[0:len(a)-dels]), c...)

	//make sure we wait until the higher partition is sorted
	wg.Wait()

	//append the sorted higher partition to the sorted primary slice
	a = append(a, b...)

	return a
}
