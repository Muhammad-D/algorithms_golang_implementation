package mergesort

import (
	"fmt"

	"github.com/Muhammad-D/algorithms_golang_implementation/internal/model"
)

func Start() {
	var size int
	fmt.Print("What length a slice should be of? ")
	fmt.Scan(&size)
	slice := model.GenerateSlice(size)
	fmt.Printf("\n-------------------------Unsorted Slice-------------------------\n\n%v\n", slice)
	slice = mergeSort(slice)
	fmt.Printf("\n--------------------------Sorted Slice--------------------------\n\n%v\n", slice)

}

func mergeSort(slice []int) []int {
	size := len(slice)
	if size == 1 {
		return slice
	}
	middle := size / 2
	left := make([]int, middle)
	right := make([]int, size-middle)

	for i, v := range slice {
		if i < middle {
			left[i] = v
		} else {
			right[i-middle] = v
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {

	slice := make([]int, len(left)+len(right))
	var i int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			slice[i] = left[0]
			left = left[1:]
			i++
		} else {
			slice[i] = right[0]
			right = right[1:]
			i++
		}
	}
	for _, v := range left {
		slice[i] = v
		i++
	}
	for _, v := range right {
		slice[i] = v
		i++
	}

	return slice
}
