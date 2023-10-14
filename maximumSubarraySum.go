package main

import "fmt"

func maxSubArraySum(arr []int) int {
	maxEndingHere := arr[0]
	maxSoFar := arr[0]

	for i := 1; i < len(arr); i++ {
		maxEndingHere = max(arr[i], maxEndingHere+arr[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum subarray sum:", maxSubArraySum(arr))
}