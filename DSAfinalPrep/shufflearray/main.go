package main

import "fmt"

func main()  {
   nums := []int{1,2,3,4,5,6,7,8,9,10}
   fmt.Println("original array ",nums)
   fmt.Println("shuffled array ",shuffle(nums))
}

func shuffle(nums []int) []int {
    newarray := make([]int,0)
	n := (len(nums))/2

    for i:=0;i<n;i++{
        newarray = append(newarray,nums[i])
        newarray = append(newarray,nums[n+i])
        
    }
    return newarray
}