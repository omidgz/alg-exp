package main

import (
	"fmt"
	"math/rand"
	"time"
)

// QuickSort algorithm
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1

	// Choosing a pivot (here the middle element)
	pivot := len(arr) / 2

	// Partitioning
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	// Recursively sort left and right parts
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

// MergeSort algorithm
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Helper function to create a random dataset
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000) // Random numbers between 0 and 9999
	}
	return slice
}

func main() {
	// Generate a random dataset of size 100000
	datasetSize := 1000000
	data := generateRandomSlice(datasetSize)

	// Make a copy of the data for each sorting algorithm
	quickSortData := make([]int, len(data))
	mergeSortData := make([]int, len(data))
	copy(quickSortData, data)
	copy(mergeSortData, data)

	// Time QuickSort
	start := time.Now()
	QuickSort(quickSortData)
	quickSortDuration := time.Since(start)

	// Time MergeSort
	start = time.Now()
	mergeSortData = MergeSort(mergeSortData) // MergeSort returns a new sorted slice
	mergeSortDuration := time.Since(start)

	// Output the results
	fmt.Printf("QuickSort took: %v\n", quickSortDuration)
	fmt.Printf("MergeSort took: %v\n", mergeSortDuration)
}
