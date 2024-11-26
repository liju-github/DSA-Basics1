package main

import "fmt"

//beginning of the list

type linkedlist struct {
	head *node
}

type node struct {
	value int
	next  *node
}

func main() {
    l := &linkedlist{}
	l.AddAtEnd(1)
	l.AddAtEnd(2)
	l.AddAtEnd(3)
	l.AddAtEnd(4)
	fmt.Println("original array")
	l.Display()
	// l.AddToBeginning(2222)
	// l.AddAfterElement(2222,555)
	// l.DeleteAtBeginning()
	// l.DeleteAtEnd()
	l.DeleteThisTarget(3)
	l.Display()
}

func (l *linkedlist) AddAtEnd(value int) {
	iterator := l.head
	newnode := &node{
		value: value,
	}
	if iterator == nil {
		l.head = newnode
		return
	}

	for ;iterator.next != nil; iterator = iterator.next {
	}

	iterator.next = newnode
}

func (l *linkedlist) Display ()  {
   iterator := l.head

   for iterator != nil{
	fmt.Printf("%v, ",iterator.value)
    iterator = iterator.next
   }
   fmt.Println()
}


func (l *linkedlist)AddToBeginning(value int)  {
	iterator := l.head
	newnode := &node{
		value: value,
		next: iterator,
	}
	
	l.head = newnode
}


func (l *linkedlist)AddAfterElement(target int,value int) {
	iterator := l.head
	newnode := &node{
		value: value,
	}

	for iterator!=nil{
		if iterator.value == target{
			temp := iterator.next
		    iterator.next = newnode
			newnode.next = temp
			return
		}
		iterator = iterator.next
	}

	fmt.Println("target value is not present")
}

func (l *linkedlist)DeleteAtBeginning()  {
	fmt.Printf("deleting %v \n",l.head.value)
	l.head = l.head.next
}

func (l *linkedlist)DeleteAtEnd()  {
	iterator := l.head

	if l.head == nil { //o(1)
        return
    }

    if l.head.next == nil { //o(1)
        l.head = nil
        return
    }

	for iterator.next != nil{ //o(n)
		if iterator.next.next == nil{
           iterator.next = nil
		   return
		}
		iterator = iterator.next
	}
}

func (l *linkedlist)DeleteThisTarget(target int)  {
	iterator := l.head

	for iterator!=nil{
        if iterator.next.value == target{
            iterator.next = iterator.next.next
			return
		}
		iterator = iterator.next
	}
	fmt.Println("target not found")
}