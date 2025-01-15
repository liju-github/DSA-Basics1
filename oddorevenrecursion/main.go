package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintOdd(arr,0)
	fmt.Println()

	PrintEven(arr, 0)

}

func PrintOdd(arr []int, current int) {
	if arr[current]%2 != 0 {
		fmt.Println(arr[current])
	}
	if current == len(arr)-1 {
		return
	}

	PrintOdd(arr, current+1)
}

func PrintEven(arr []int, current int) {

	if arr[current]%2 == 0 {
		fmt.Println(arr[current])
	}
	if current == len(arr)-1 {
		return
	}

	PrintEven(arr, current+1)
}
