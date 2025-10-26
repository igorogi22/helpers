package arrays

import (
	"math/rand"
)

func CreateNumericArray(size int, start ...int) []int {
	s := 1

	if len(start) > 0 {
		s = start[0]
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = s + i
	}
	return arr
}

func GetRandomItem[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

func Reverse[T any](arr []T) []T {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}

func ShuffleArray[T any](arr []T) []T {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
