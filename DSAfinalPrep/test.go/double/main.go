package main

import "fmt"

type linkedlist struct {
	head *node
	tail *node
}

type node struct {
	value int
	prev  *node
	next  *node
}

func main() {
	l := &linkedlist{}

	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)
	l.InsertBeginning(4)
	l.Display()
	fmt.Println("\nafter inserting:")
	// l.InsertBeforeTarget(4,555)
	// l.InsertAfterTarget(2,555)
	// l.InsertEnd(22)
	// l.DeleteAtBeginning()
	// l.DeleteAtEnd()
	// l.DeleteTargetNode(2)
	l.Display()
}

func (l *linkedlist) InsertBeginning(value int) {
	newnode := &node{
		value: value,
	}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
		return
	}

	newnode.next = l.head
	l.head.prev = newnode
	l.head = newnode 
}

func (l *linkedlist) Display() {
	if l.head == nil {
		return
	}

	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}


func (l *linkedlist)InsertAfterTarget(target int,value int)  {
	if l.head == nil{
		return
	}

	current := l.head // // //
	for current!= nil{
		if current.value == target{
			newnode := &node{
               value: value,
			   prev: current,
			   next: current.next,
			}

			current.next.prev = newnode
			current.next = newnode
			return
		}
		current = current.next
		if current.next == nil{
			l.tail = current
			return
		}
	}


}

func (l *linkedlist) InsertBeforeTarget(target int, value int) {
    if l.head == nil {
        return
    }

    current := l.head
    for current != nil {
        if current.value == target { 
            newnode := &node{
                value: value,
                next:  current,
                prev:  current.prev, 
            }


            if current.prev != nil {
                current.prev.next = newnode 
            } else { 
                l.head = newnode
            }

            current.prev = newnode 
            newnode.next = current 
            return
        }
        current = current.next
		if current.next == nil{
			l.tail = current
			return
		}
    }
}

func (l *linkedlist)InsertEnd(value int)  {
	newnode := &node{
		value: value,
	}
	if l.tail == nil{
		l.tail = newnode
		l.head = newnode
		return
	}

	newnode.prev = l.tail
	l.tail.next = newnode
	l.tail = newnode
    
}

func (l *linkedlist)DeleteAtBeginning()  {
	if l.head != nil{
		l.head = l.head.next
		l.head.prev = nil
	}


}

func (l *linkedlist)DeleteAtEnd()  {
	if l.head == nil{
		return
	}
	if l.head == l.tail{
		l.head = nil
		l.tail =nil
		return
	}

	l.tail.prev.next = nil
}

func (l *linkedlist)DeleteTargetNode(target int)  {
	//if target == l.head
	   //if l.head.next == nil
	//if target == l.tail
	   //l.tail.prev = nil
    
	if l.head == nil{
		return
	}

	if l.head.value == target{
		if l.head.next==nil{
			l.head = nil
			return
		}
		l.head = l.head.next
		l.head.prev = nil
	}
    
	if l.tail.value ==  target{
	    l.tail = l.tail.prev
		l.tail.next = nil
	}

	current := l.head
	for current.next !=nil{

		if current.next.value == target{
			current.next = current.next.next
			current.next.prev = current
			return
		}


		current = current.next
		if current.next == nil{
			l.tail = current
			return
		}

	}
  
}
