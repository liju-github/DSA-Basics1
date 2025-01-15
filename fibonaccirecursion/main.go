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

// package main

// import ("fmt")

// func main(){
//     limit := 10;
//     arr := make([]int,limit);
    
//     fmt.Println("fibonacci : ", fibonacci(arr,0,limit));
// }

// func fibonacci(arr []int,current int , limit int)([]int){
//     if current >= limit{
//         return arr
//     }
    
//     if current == 0{
//         arr[0]= 0
//     }else if current == 1{
//         arr[1] = 1
//     }else{
//         arr[current] = arr[current-1] + arr[current-2]
//     }
//     return fibonacci(arr,current+1,limit)
// }