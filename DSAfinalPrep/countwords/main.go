package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "a boy with an umbrella, an a boy with"
    fmt.Println(CountWords(str))
}

func CountWords(str string) map[string]int {
    str = strings.ToLower(str)
    
	str = strings.ReplaceAll(str, ",", "")
    
    words := strings.Fields(str)
    
    count := make(map[string]int)
    
    for _, word := range words {
        count[word]++
    }
    
    return count
}
