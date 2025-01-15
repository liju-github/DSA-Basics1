package main

import "fmt"

//remove the nth node from the end
type linkedlist struct{
	head *node
}

type node struct{
	value int
	next *node
}

func main(){
	l := &linkedlist{}
	l.AddAtBeginning(1)
	l.AddAtBeginning(2)
	l.AddAtBeginning(3)
	l.AddAtBeginning(4)
	l.AddAtBeginning(5)

	l.Display()

	//delete the 2'nd last node n=2 ie. node.value = 5
	l.Delete(2)

	l.Display()
}

func (l *linkedlist)Delete(n int)  {
	dummy := l.head
	slow,fast := dummy,dummy

	for i:=0;i<n+1;i++{
		fmt.Println("loop running") 
		fast = fast.next
	} 

	for fast!=nil{
		fmt.Println("iter running")
		fast = fast.next 
		slow = slow.next 
	}

	slow.next = slow.next.next 


}

func (l *linkedlist)AddAtBeginning(value int)  {
	newnode:= &node{
		value: value,
	}

	newnode.next = l.head
	l.head = newnode
}

func (l *linkedlist)Display()  {
	iter := l.head

	for iter !=nil{
		fmt.Print("->",iter.value)
		iter = iter.next
	}
	fmt.Println()
}