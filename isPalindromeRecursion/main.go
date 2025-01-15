package main

import (
	"fmt"
)

func main(){

	str:="2323dsdsd999"
	fmt.Println(isPalindromeRecursion(str,0,len(str)-1))

}


func isPalindromeRecursion(str string,left,right int) bool {
	r := []rune(str)

	if left >= right {
		return true
	}

	if r[left]!=r[right]{
		return false
		
	}

	return isPalindromeRecursion(str,left+1,right-1)


}

// package main

// import ("fmt")

// func main(){
//     str:="malayalam";
    
//     fmt.Println("palindrome check :",palindrome(str,0,len(str)-1))
// }

// func palindrome(str string,left int,right int) bool{
//     if left > right{
//         return true
//     }
    
//     if str[left] != str[right]{
//         return false
//     }
//     return palindrome(str,left+1,right-1)
// }