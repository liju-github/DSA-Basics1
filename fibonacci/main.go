package main

import "fmt"

func main() {
    result := fibonacci(10)
    
    for _, v := range result {
        fmt.Println(v)
    }
}

func fibonacci(limit int) []int {

    arr := make([]int, limit)

    if limit > 0 {
        arr[0] = 0 
    }
    if limit > 1 {
        arr[1] = 1 
    }

    for i := 2; i < limit; i++ {
        arr[i] = arr[i-1] + arr[i-2]
    }

    return arr
}


// package main
// import "fmt"

// func main() {
//   limit := 10
//   fmt.Println("fibonacci : ",fibonacci(limit))
// }

// func fibonacci(limit int)[]int{
//     arr := make([]int,limit)
//     arr[0],arr[1] = 0,1
    
//     for i:=2;i<limit;i++{
//         arr[i] = arr[i-1]+arr[i-2];
//     }
    
//     return arr
// }