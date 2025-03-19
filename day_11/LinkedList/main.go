package main

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Node(data T) *Node[T] {
	return &Node[T]{data: data, next: nil}
}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current.next != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println(current.data)
}

func (l *LinkedList[T]) Size() int {
	size := 0
	current := l.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (l *LinkedList[T]) InsertAtTail(data T) {
	if l.head == nil {
		l.head = &Node[T]{data: data, next: nil}
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node[T]{data: data, next: nil}

}

func (l *LinkedList[T]) InsertAtPos(data T, pos int) {
	if pos <= 0 {
		fmt.Println("Invalid Position!")
		return
	}
	if l.head == nil {
		l.head = &Node[T]{data: data, next: nil}
		return
	}
	if pos == 1 {
		l.InsertAtHead(data)
		return
	}
	if pos >= l.Size() {
		l.InsertAtTail(data)
		return
	}
	// fmt.Println("here")

	current := l.head
	for i := 1; i <= pos-1; i++ {
		current = current.next
	}

	newNode := &Node[T]{data: data, next: nil}
	newNode.next = current.next
	current.next = newNode

}

func (l *LinkedList[T]) InsertAtHead(data T) {
	if l.head == nil {
		l.head = &Node[T]{data: data, next: nil}
		return
	}
	newNode := &Node[T]{data: data, next: nil}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) DeleteFromHead() {
	if l.head == nil {
		panic("No Node to delete!")
	}
	if l.head.next == nil {
		l.head = nil
		return
	}
	l.head = l.head.next
}

func (l *LinkedList[T]) DeleteFromTail() {
	if l.head == nil {
		fmt.Println("No node to delete!")
	}
	if l.head.next == nil {
		l.head = nil
		return
	}
	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (l *LinkedList[T]) DeleteFromPos(pos int) {
	if pos <= 0 {
		fmt.Printf("Invalid position: %v\n", pos)
		return
	}
	if pos > l.Size() {
		fmt.Printf("postion out of range!: %v\n", pos)
		return
	}

	if pos == 1 {
		l.DeleteFromHead()
		return
	}

	if pos == l.Size() {
		l.DeleteFromTail()
		return
	}

	current := l.head
	cur_pos := 1
	for cur_pos != pos-1 && current != nil {
		current = current.next
		cur_pos++
	}

	current = current.next.next
}

func main() {
	LL := LinkedList[int]{}
	LL.InsertAtTail(1)
	LL.InsertAtTail(2)
	LL.InsertAtTail(4)
	LL.InsertAtHead(0)
	LL.InsertAtPos(3, 3)
	LL.Print()
	LL.DeleteFromHead()
	LL.DeleteFromPos(0)
	LL.Print()
}
