package main

import (
	"fmt"
)

func main(){

	// str := " hello world 1211 ??/"
	str:="malayalam"
	fmt.Println(isPalindrome(str))


}


func isPalindrome(str string)bool  {
	
	r := []rune(str)

	left,right := 0,len(str)-1

	for left < right{
		if r[left]!=r[right]{
			return false
		}
		left++
		right--
	}

	return true
}
