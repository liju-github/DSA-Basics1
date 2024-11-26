package main

import "fmt"

// linkedlist struct
type linkedlist struct {
	mainhead *node
}

// node struct
type node struct {
	value int
	next  *node
}

// main function
func main() {
	l := &linkedlist{}
	l.Add(3)
	l.Add(5)
	l.Add(7)
	l.Display()
}

func (l *linkedlist) Add(value int) {
	newnode := &node{
		value: value,
	}
	if l.mainhead == nil {
		l.mainhead = newnode
		return
	}

	current := l.mainhead
	for current.next != nil {
		current = current.next
	}
	current.next = newnode
}

func (l *linkedlist) Display() {
	current := l.mainhead
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
