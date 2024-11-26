package main

import "fmt"

type node struct {
	value int
	tail  *node
}

type linkedlist struct {
	mainhead *node
}

func main() {
	l := &linkedlist{}
	l.AddAtEnd(1)
	l.AddAtEnd(1)
	l.AddAtEnd(2)
	l.AddAtEnd(3)
	l.AddAtEnd(4)
	l.AddAtEnd(5)
	l.Display()
	l.DeleteEven()
	l.Display()
	l.DeleteOdd()
	l.Display()
}

func (l *linkedlist) AddAtEnd(value int) {
	iterator := l.mainhead
	newnode := &node{
		value: value,
	}
	if iterator == nil {
		l.mainhead = newnode
		return
	}

	for ; iterator.tail != nil; iterator = iterator.tail {
	}

	iterator.tail = newnode
}

func (l *linkedlist) Display() {
	iterator := l.mainhead

	for iterator != nil {
		fmt.Printf("%v ", iterator.value)
		iterator = iterator.tail
	}
	fmt.Println("\n")
}

func (l *linkedlist) DeleteOdd() {
	if l.mainhead == nil {
		return
	}

	for l.mainhead != nil && l.mainhead.value%2 != 0 {
		l.mainhead = l.mainhead.tail
	}

	current := l.mainhead

	for current != nil && current.tail != nil {
		if current.tail.value%2 != 0 {
			current.tail = current.tail.tail
		} else {
			current = current.tail
		}
	}
}

func (l *linkedlist) DeleteEven() {
	if l.mainhead == nil {
		return
	}

	for l.mainhead != nil && l.mainhead.value%2 == 0 {
		l.mainhead = l.mainhead.tail
	}

	current := l.mainhead

	for current != nil && current.tail != nil {
		if current.tail.value%2 == 0 {
			current.tail = current.tail.tail
		} else {
			current = current.tail
		}
	}
}