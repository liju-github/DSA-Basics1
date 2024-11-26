package main

import "fmt"



func main()  {
	arr:=[]int{1,2,3,4,5,6,7,8,9,10}
	max,min := MaximumMinimum(arr)
	fmt.Printf("Maximum is %v and Minimum is %v \n",max,min)
}


func MaximumMinimum(arr []int) (int,int)  {
	max,min := arr[0],arr[0]
	for _,v := range arr{
		if v > max{
			max =v
		}

	    if v < min{
			min=v
		}
	}

	return max,min
}