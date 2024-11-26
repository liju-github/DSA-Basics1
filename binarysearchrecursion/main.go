package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4, 6, 3, 9, 10, 23, 67, 40}
	sort.Ints(arr)

	index := binarySearchRecursive(arr, 67, 0, len(arr)-1)
	if index != -1 {
		fmt.Printf("%v value found at %v index \n", arr[index], index)
	} else {
		fmt.Println("not found")
	}
}

func binarySearchRecursive(arr []int, target int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}	
