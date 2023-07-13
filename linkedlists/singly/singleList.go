package singly

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	size int64
}

func (list *SinglyLinkedList) InsertAtLast(Data interface{}) {
	newNode := &Node{
		Data: Data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}

	list.size++
}

func (list *SinglyLinkedList) Size() int64 {
	return list.size
}

func (list *SinglyLinkedList) InsertAtStart(Data interface{}) {
	newNode := &Node{
		Data: Data,
		Next: list.Head,
	}

	list.Head = newNode
	if list.Tail == nil {
		list.Tail = newNode
	}

	list.size++
}

func (list *SinglyLinkedList) DeleteElement(Data interface{}) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == Data {
		list.Head = list.Head.Next
		if list.Head == nil {
			list.Tail = nil
		}
		list.size--
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == Data {
			if current.Next == list.Tail {
				list.Tail = current
			}
			current.Next = current.Next.Next
			list.size--
			return
		}
		current = current.Next
	}
}

func (list *SinglyLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
