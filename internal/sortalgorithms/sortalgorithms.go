package sortalgorithms

import (
	"fmt"

	"github.com/Muhammad-D/algorithms_golang_implementation/internal/model"
)

func Start() {
	var size int
	fmt.Print("How long slice should be: ")
	fmt.Scan(&size)

	slice := model.GenerateSlice(size)
	fmt.Printf("\n--- Unsorted List ---\n\n%v\n", slice)

	fmt.Printf("\n--- Sorted List ---\n\n%v\n", mergeSort(slice))

}

func mergeSort(slice []int) []int {

	size := len(slice)

	if size == 1 {
		return slice
	}

	middle := int(size / 2)

	var (
		left  = make([]int, middle)
		right = make([]int, size-middle)
	)

	for i, sliceValue := range slice {

		if i < middle {
			left[i] = sliceValue
		} else {
			right[i-middle] = sliceValue
		}

	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {

	result = make([]int, len(left)+len(right))
	var i int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for _, v := range left {
		result[i] = v
		i++
	}
	for _, v := range right {
		result[i] = v
		i++
	}

	return
}
