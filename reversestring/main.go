package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world ! !"
	fmt.Println(reversestring(str))
}

func reversestring(str string) string {

	words := strings.Split(str,"")
	
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	
	return strings.Join(words, "")
}


// package main

// import ("fmt")

// func main(){
//     str := "helloworld"
//     strarr := []byte(str)
//     fmt.Println("reverse : ",reverse(strarr,0,len(str)-1))
// }

// func reverse(strarr []byte,left,right int) string{
//     if left > right{
//         return string(strarr)
//     }
    
//     strarr[left],strarr[right] = strarr[right],strarr[left]
    
//     return reverse(strarr,left+1,right-1)
// }