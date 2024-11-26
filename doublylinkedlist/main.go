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
    newNode := &node{
        prev:  nil,
        next:  l.mainhead,
        value: value,
    }

    if l.mainhead != nil {
        l.mainhead.prev = newNode
    }

    l.mainhead = newNode
}
func (l *linkedlist) Display() {
	iterator := l.mainhead

	for iterator != nil {
		// fmt.Printf(" Value : %v  || Previous : %v ||  Next : %v",iterator.value,iterator.prev,iterator.next)
		fmt.Printf(" %v ",iterator.value)
		iterator = iterator.next
	}
	fmt.Printf("\n")
}


func (l *linkedlist)InsertAtEnd(value int)  {
	current := l.mainhead
 
 	for current != nil{
		if current.next == nil{
			newnode := &node{
				prev: current,
				value:value,
				next: nil,
			}

			current.next = newnode
			return
		}
       current = current.next
	}
}

func (l *linkedlist)InsertAfterTarget(target int,value int)  {
	current := l.mainhead

	for current !=nil{
		if current.value == target{
			newnode := &node{
				prev: current,
				value: value,
				next: current.next,
			}

			current.next = newnode;return
		}
       current = current.next
	}
}

func (l *linkedlist)InsertBeforeTarget(target int,value int)  {
	current := l.mainhead  //1 //new //2 //3

	for current !=nil{
	    if current.next.value == target{
            newnode := &node{
				prev: current,
				value: value,
				next: current.next,
			}

			current.next = newnode;return

		} 
		current = current.next

	}
}


func (l *linkedlist)DeleteAtBeginning( )  {
	current := l.mainhead

	if current == nil{
		return
	}

	current.next.prev = nil
	l.mainhead = current.next 
}

func  (l *linkedlist)DeleteAtEnd()  {
	current := l.mainhead

	if current == nil || current.next == nil{
		current = nil
		return
	}

	for current != nil{
		if current.next.next == nil{
            current.next =nil
			return
		}
		current = current.next
	}
}

func (l *linkedlist) DeleteThisTarget(target int) {
    current := l.mainhead
    if current == nil {
        return
    }
    if current.value == target {
        l.mainhead = current.next
    }
    
    for current != nil {
        if current.value == target {
            if current.next == nil && current.prev == nil {
                l.mainhead = nil
            } else if current.prev == nil {
                l.mainhead = current.next
                current.next.prev = nil
            } else if current.next == nil {
                current.prev.next = nil
            } else {
                current.prev.next = current.next
                current.next.prev = current.prev
            }
        }
        current = current.next
    }
}