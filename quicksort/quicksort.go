package quicksort

import "math/rand"

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	left := 0
	right := len(a) - 1
	pivot := rand.Int() % len(a)

	temp := a[pivot]
	a[pivot] = a[right]
	a[right] = temp

	for i, _ := range a {
		if a[i] < a[right] {
			c := a[left]
			a[left] = a[i]
			a[i] = c
			left++
		}
	}

	temp = a[left]
	a[left] = a[right]
	a[right] = temp

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}
