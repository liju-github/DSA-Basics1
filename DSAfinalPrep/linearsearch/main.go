package main

import "fmt"



func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	index := LinearSearch(arr,3)
	if index != -1{
		fmt.Println("value found at index ",index)
	}else{
		fmt.Println("value not found")
	}
}


func LinearSearch(arr []int,target int) int {
	for i,v := range arr{
		if v == target{
			return i
		}
	}

	return -1
}