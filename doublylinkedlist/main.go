package main

import "fmt"

type node struct {
	prev  *node
	value int
	next  *node
}

type linkedlist struct {
	mainhead *node
}

// Constructor for creating a new node
func NewNode(prev *node, value int, next *node) *node {
	return &node{
		prev:  prev,
		value: value,
		next:  next,
	}
}

func main() {
	l := &linkedlist{}
	l.InsertAtBeginning(1)
	l.InsertAtBeginning(2)
	l.InsertAtBeginning(3)
	l.InsertAtBeginning(4)
	l.InsertAtBeginning(5)
	fmt.Println("Original list:")
	l.Display()

	l.DeleteThisTarget(2)
	fmt.Println("After deleting 2:")
	l.Display()
}

func (l *linkedlist) InsertAtBeginning(value int) {
	newNode := NewNode(nil, value, l.mainhead)

	if l.mainhead != nil {
		l.mainhead.prev = newNode
	}

	l.mainhead = newNode
}

func (l *linkedlist) Display() {
	iterator := l.mainhead

	for iterator != nil {
		fmt.Printf(" %v ", iterator.value)
		iterator = iterator.next
	}
	fmt.Println()
}

func (l *linkedlist) InsertAtEnd(value int) {
	if l.mainhead == nil {
		l.InsertAtBeginning(value)
		return
	}

	current := l.mainhead
	for current.next != nil {
		current = current.next
	}

	newNode := NewNode(current, value, nil)
	current.next = newNode
}

func (l *linkedlist) InsertAfterTarget(target int, value int) {
	current := l.mainhead

	for current != nil {
		if current.value == target {
			newNode := NewNode(current, value, current.next)
			if current.next != nil {
				current.next.prev = newNode
			}
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (l *linkedlist) InsertBeforeTarget(target int, value int) {
	current := l.mainhead

	if current != nil && current.value == target {
		l.InsertAtBeginning(value)
		return
	}

	for current != nil {
		if current.next != nil && current.next.value == target {
			newNode := NewNode(current, value, current.next)
			current.next.prev = newNode
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (l *linkedlist) DeleteAtBeginning() {
	if l.mainhead == nil {
		return
	}

	if l.mainhead.next != nil {
		l.mainhead.next.prev = nil
	}

	l.mainhead = l.mainhead.next
}

func (l *linkedlist) DeleteAtEnd() {
	if l.mainhead == nil || l.mainhead.next == nil {
		l.mainhead = nil
		return
	}

	current := l.mainhead
	for current.next != nil {
		current = current.next
	}

	current.prev.next = nil
}

func (l *linkedlist) DeleteThisTarget(target int) {
	current := l.mainhead

	for current != nil {
		if current.value == target {
			if current.prev == nil { // Deleting the head node
				l.mainhead = current.next
				if current.next != nil {
					current.next.prev = nil
				}
			} else if current.next == nil { // Deleting the tail node
				current.prev.next = nil
			} else { // Deleting a middle node
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			return
		}
		current = current.next
	}
}
