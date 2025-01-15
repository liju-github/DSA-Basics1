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


// package main


// import (
//     "fmt"
//     "sort"
//     )
    
// func main(){
//     arr := []int{2,3,7,1,9,10,12,66,22,15}
//     sort.Ints(arr)
//     fmt.Println("Sorted Array: ",arr)
//     target := 22
//     index := BinarySearch(arr,target,0,len(arr)-1)
//     if index == -1{
//         fmt.Println("The target is not found in the arr")
//     }else{
//         fmt.Printf("The target: %v is found at index %v \n",target,index)
//     }
// }

// func BinarySearch(arr []int,target int, left int, right int)int{
//     if left > right{
//         return -1
//     }
    
//     mid := (left+right)/2
    
//     if target == arr[mid]{
//         return mid
//     }else if target > arr[mid]{
//         return BinarySearch(arr,target,mid+1,right)
//     }else{
//         return BinarySearch(arr,target,left,mid-1)
//     }
    
//     return -1
// }

