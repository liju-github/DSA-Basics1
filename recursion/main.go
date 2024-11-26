package main

import "fmt"

func main() {
	count(5)
}

func count(n int) {
	if n >= 0 {
		fmt.Println(n)
	} else {
		return
	}
	count(n - 1)
}
