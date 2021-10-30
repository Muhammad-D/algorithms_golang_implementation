package model

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {

	mySlice := make([]int, size, size)

	rand.Seed(time.Now().UnixNano())

	for i := range mySlice {
		mySlice[i] = rand.Intn(999) - rand.Intn(999)
	}

	return mySlice

}
