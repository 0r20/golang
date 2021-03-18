package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type linkedList struct {
	head   *node
	foot   *node
	length int
}

func (l *linkedList) PushFront(n *node) {
	if l.length == 0 {
		l.foot = n
	} else {
		l.head.prev = n
	}
	prevHead := l.head
	l.head = n
	l.head.next = prevHead
	l.length++
}

func (l *linkedList) PushBack(n *node) {
	if l.length == 0 {
		l.head = n
	} else {
		l.foot.next = n
	}
	prevFoot := l.foot
	l.foot = n
	l.foot.prev = prevFoot
	l.length++
}

func (l *linkedList) PrintFromHead() {
	for pointer := l.head; pointer != nil; pointer = pointer.next {
		fmt.Printf("%d ", pointer.data)
	}
	fmt.Printf("\n")
}

func (l *linkedList) PrintFromFoot() {
	for pointer := l.foot; pointer != nil; pointer = pointer.prev {
		fmt.Printf("%d ", pointer.data)
	}
	fmt.Printf("\n")
}

func (l *linkedList) Remove(item int) {
	if l.length == 0 {
		return
	}
	if l.head.data == item {
		l.head = l.head.next
		l.length--
		return
	}
	prevPointer := l.head
	for prevPointer.next.data != item {
		if prevPointer.next.next == nil {
			return
		}
		prevPointer = prevPointer.next
	}
	prevPointer.next = prevPointer.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	myList.PushFront(node1)
	myList.PushFront(node2)
	myList.PushFront(node3)
	myList.PushBack(node4)
	myList.PrintFromHead()
	// myList.PrintFromFoot()
	// myList.Remove(100)
	// myList.print()
	// fmt.Printf("%d\n", myList.head.data)
}
