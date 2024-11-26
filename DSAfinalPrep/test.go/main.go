package main

import "fmt"


type linkedlist struct{
	mainhead *node
}

type node struct{
	value int
	next *node
}

func main()  {
	l := &linkedlist{}
	l.AddBeginning(1)
	l.AddBeginning(2)
	l.AddBeginning(3)
	l.Display()
	fmt.Println()
	// l.AddAtEnd(555)
	// l.InsertAfterTarget(555,777)
	// l.DeleteAtBeginning()
	// l.DeleteAtEnd()
	l.DeleteTargetNode(2)

	l.Display()

}


func (l *linkedlist)AddBeginning(value int)  {

	newnode := &node{
		value: value,
	}

	if l.mainhead == nil{
        l.mainhead = newnode
		return
	}

    newnode.next = l.mainhead
	l.mainhead = newnode

}


func (l *linkedlist)Display()  {
	current := l.mainhead

	for current!= nil{
		fmt.Println(current.value)
		current = current.next
	}
}

func (l *linkedlist)AddAtEnd(value int)  {

	newnode := &node{
		value: value,
	}

	if l.mainhead == nil{
		l.mainhead = newnode
		return
	}

	current := l.mainhead
	for current.next != nil{
		current = current.next
	}

	current.next = newnode

}

func (l *linkedlist) InsertAfterTarget(target int, value int) {
    if l.mainhead == nil {
        return
    }

    newNode := &node{
        value: value,
    }

    if l.mainhead.value == target {
        newNode.next = l.mainhead.next
        l.mainhead.next = newNode
        return
    }

    current := l.mainhead
    for current.next != nil {
        if current.value == target {
            newNode.next = current.next
            current.next = newNode
            return
        }
        current = current.next
    }

	current.next  =newNode
}

func (l *linkedlist)DeleteAtBeginning()  {
	if l.mainhead == nil{
		return
	}

	l.mainhead = l.mainhead.next
}

func (l *linkedlist)DeleteAtEnd()  {
	if l.mainhead == nil || l.mainhead.next == nil{
		l.mainhead = nil
		return
	}

	current := l.mainhead

	for current.next.next != nil{
       current = current.next
	}

	
	current.next = nil
}

func (l *linkedlist) DeleteTargetNode(target int) {
    if l.mainhead == nil {
        return 
    }

    if l.mainhead.value == target {
        l.mainhead = l.mainhead.next
        return
    }

    current := l.mainhead
    for current.next != nil {
        if current.next.value == target {
            current.next = current.next.next
            return
        }
        current = current.next
    }
}