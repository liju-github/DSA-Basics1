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

