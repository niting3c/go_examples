package singly

import "fmt"

type SingleNode struct {
	Data interface{}
	Next *SingleNode
}

type SinglyLinkedList struct {
	Head *SingleNode
	Tail *SingleNode
	size int64
}

func (list *SinglyLinkedList) InsertAtLast(Data interface{}) {
	newSingleNode := &SingleNode{
		Data: Data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newSingleNode
		list.Tail = newSingleNode
	} else {
		list.Tail.Next = newSingleNode
		list.Tail = newSingleNode
	}

	list.size++
}

func (list *SinglyLinkedList) Size() int64 {
	return list.size
}

func (list *SinglyLinkedList) InsertAtStart(Data interface{}) {
	newSingleNode := &SingleNode{
		Data: Data,
		Next: list.Head,
	}

	list.Head = newSingleNode
	if list.Tail == nil {
		list.Tail = newSingleNode
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
	fmt.Print("element not found , not deleting")
}

func (list *SinglyLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (list *SinglyLinkedList) InsertAfter(nodeElement interface{}, data interface{}) {
	current := list.Head
	node := &SingleNode{
		Data: data,
	}
	for current != nil {
		if current.Data == nodeElement {
			node.Next = current.Next
			current.Next = node
			if node.Next == nil {
				list.Tail = node
			}
			list.size++
			return
		}
		current = current.Next
	}
	fmt.Print("element not found , not inserting")
}
