package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    shift := 1

    fmt.Println("Original array:", arr)
    fmt.Println("Shifted array:", ShiftElements(arr, shift))
}

func ShiftElements(arr []int, shift int) []int {
    n := len(arr)
    shift = shift % n
    newArray := make([]int, n)
    for i, v := range arr {
        newPos := (i + shift) % n
        newArray[newPos] = v
    }

    return newArray
}
