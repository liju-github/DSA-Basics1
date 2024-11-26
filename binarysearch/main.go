package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1,2,3,4,5,6,34,3,4,34,3,4,2,3,2,30}
	
	sort.Ints(arr)
	fmt.Println(arr)
	
	binarysearch(arr, 30)
}

func binarysearch(arr []int, find int) {
	// startTime := time.Now()

	left, right := 0, len(arr)-1
	
	for left <= right {
		fmt.Printf("left = %v right = %v \n",left,right)
		mid := left + (right-left)/2
		if arr[mid] == find {
			fmt.Println("the number is found at", mid)
			return
		} else if arr[mid] < find {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println("Element not found")

	// elapsedTime := time.Since(startTime).Seconds()
	// fmt.Printf("Time taken: %v seconds\n", elapsedTime)
}