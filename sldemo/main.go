package main

import "fmt"

type linkedlist struct {
	head *node
}

type node struct {
	next  *node
	value int
}

func main() {
	l := &linkedlist{}

	l.AddAtBeginning(1)
	l.AddAtBeginning(2)
	l.AddAtBeginning(3)
	l.AddAtBeginning(4)
	l.AddAtBeginning(5)

	l.AddAtEnd(100)

	l.AddAfterTarget(300,100)

	l.Display()

	l.DeleteAtBeginning()

	l.Display()

	l.DeleteAtEnd()
	l.Display()

	l.DeleteTarget(100)
	l.Display()


}

func (l *linkedlist)DeleteTarget(target int){
	if l.head == nil{
		return;
	}

	iter := l.head;

	for iter.next!=nil{
		if iter.next.value == target {
			if iter.next.next !=nil{
				iter.next = iter.next.next
				return
			}else{
				iter.next = nil
				return
			}
		}

		iter = iter.next
	}
}

func (l *linkedlist)DeleteAtEnd()  {
	if l.head == nil || l.head.next == nil{
		l.head = nil;return
	}
	iter := l.head
	for iter.next.next!=nil{
		iter = iter.next
	}
	iter.next = nil

}

func (l *linkedlist)DeleteAtBeginning()  {
	if l.head == nil || l.head.next ==nil{
		l.head = nil
		return
	}
	l.head = l.head.next
	
}

func (l *linkedlist)AddAfterTarget(value,target int)  {
	newnode := &node{
		value:value,
	}

	iter := l.head
	for iter!=nil{
		if iter.value == target{
			temp := iter.next
			iter.next = newnode
			newnode.next = temp
		}
		iter = iter.next
	}
}

func (l *linkedlist)AddAtEnd(value int)  {
	newnode := &node{
		value: value,
	}

	iter := l.head
	for iter.next!=nil{
		iter=iter.next	
	}

	iter.next = newnode
}

func (l *linkedlist) AddAtBeginning(value int) {
	newnode := node{
		value: value,
	}

	iter := l.head
	newnode.next = iter
	l.head = &newnode
}

func (l *linkedlist) Display() {
	iter := l.head

	for iter != nil {
		fmt.Print("-> ",iter.value)
		iter = iter.next
	}
	fmt.Println()
}
