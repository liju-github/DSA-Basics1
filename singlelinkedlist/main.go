package main

import "fmt"

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// Node represents a single node in the linked list
type Node struct {
	value int
	next  *Node
}

func main() {
	l := &LinkedList{}
	l.AddAtEnd(1)
	l.AddAtEnd(2)
	l.AddAtEnd(3)
	l.AddAtEnd(4)
	fmt.Println("Original list:")
	l.Display()

	l.AddToBeginning(2222)
	l.AddAfterElement(2, 555)
	l.DeleteAtBeginning()
	l.DeleteAtEnd()
	l.DeleteThisTarget(3)
	fmt.Println("Updated list:")
	l.Display()
}

// AddAtEnd adds a new node with the given value at the end of the linked list
// Time Complexity: O(n)
func (l *LinkedList) AddAtEnd(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	iterator := l.head
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = newNode
}

// Display prints all the elements of the linked list
// Time Complexity: O(n)
func (l *LinkedList) Display() {
	iterator := l.head
	for iterator != nil {
		fmt.Printf("%v -> ", iterator.value)
		iterator = iterator.next
	}
	fmt.Println("nil")
}

// AddToBeginning adds a new node with the given value at the beginning of the linked list
// Time Complexity: O(1)
func (l *LinkedList) AddToBeginning(value int) {
	newNode := &Node{value: value, next: l.head}
	l.head = newNode
}

// AddAfterElement adds a new node with the given value after the first occurrence of the target value
// Time Complexity: O(n)
func (l *LinkedList) AddAfterElement(target int, value int) {
	iterator := l.head
	for iterator != nil {
		if iterator.value == target {
			newNode := &Node{value: value, next: iterator.next}
			iterator.next = newNode
			return
		}
		iterator = iterator.next
	}
	fmt.Println("Target value not found")
}

// DeleteAtBeginning removes the first node of the linked list
// Time Complexity: O(1)
func (l *LinkedList) DeleteAtBeginning() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Printf("Deleting: %v\n", l.head.value)
	l.head = l.head.next
}

// DeleteAtEnd removes the last node of the linked list
// Time Complexity: O(n)
func (l *LinkedList) DeleteAtEnd() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.next == nil {
		fmt.Printf("Deleting: %v\n", l.head.value)
		l.head = nil
		return
	}
	iterator := l.head
	for iterator.next.next != nil {
		iterator = iterator.next
	}
	fmt.Printf("Deleting: %v\n", iterator.next.value)
	iterator.next = nil
}

// DeleteThisTarget removes the first node with the given target value
// Time Complexity: O(n)
func (l *LinkedList) DeleteThisTarget(target int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.value == target {
		fmt.Printf("Deleting: %v\n", l.head.value)
		l.head = l.head.next
		return
	}
	iterator := l.head
	for iterator.next != nil {
		if iterator.next.value == target {
			fmt.Printf("Deleting: %v\n", iterator.next.value)
			iterator.next = iterator.next.next
			return
		}
		iterator = iterator.next
	}
	fmt.Println("Target not found")
}