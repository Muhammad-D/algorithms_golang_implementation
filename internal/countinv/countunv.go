package countinv

import (
	"fmt"

	"github.com/Muhammad-D/algorithms_golang_implementation/internal/model"
)

func Start() {
	mySlice := model.GenerateSlice(8)
	fmt.Println("\nUnsorted list:", mySlice, "\n ")
	inv := mergeSort(mySlice)
	fmt.Println("A number of inversions:", inv, "\n ")
	fmt.Println("Sorted list:  ", mySlice, "\n ")

}

func mergeSort(arr []int) int {

	tempSlice := make([]int, len(arr))

	return _mergeSort(arr, tempSlice, 0, len(arr))
}

func _mergeSort(arr, temp []int, left, right int) int {
	var mid, countInvr int

	if right-1 <= left {
		return countInvr
	}

	mid = (left + right - 1) / 2
	countInvr += _mergeSort(arr, temp, left, mid+1)
	countInvr += _mergeSort(arr, temp, mid+1, right)
	countInvr += merge(arr, temp, left, mid+1, right)

	return countInvr
}

func merge(arr, temp []int, left, mid, right int) int {

	var i, j, k = left, mid, left
	var countInvr int

	for i < mid && j < right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
			countInvr = countInvr + (mid - i)
		}
	}
	for ; i < mid; i, k = i+1, k+1 {
		temp[k] = arr[i]
	}
	for ; j < right; j, k = j+1, k+1 {
		temp[k] = arr[j]
	}
	for m := left; m < right; m++ {
		arr[m] = temp[m]
	}

	return countInvr
}
