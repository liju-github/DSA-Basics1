package main

import "fmt"

func main() {
	limit := 10
	arr := make([]int, limit)
	FibonacciRecursion(arr, limit, 0)
	fmt.Println(arr)
}

func FibonacciRecursion(arr []int, limit int, current int) {
	if current >= limit {
		return
	}
	if current == 0 {
		arr[current] = 0
	} else if current == 1 {
		arr[current] = 1
	} else {
		arr[current] = arr[current-1] + arr[current-2]
	}
	FibonacciRecursion(arr, limit, current+1)
}