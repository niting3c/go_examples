package doubly

import "fmt"

type DoubleNode struct {
	Data interface{}
	next *DoubleNode
	prev *DoubleNode
}

type DoublyLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	size int64
}

func (list *DoublyLinkedList) Size() int64 {
	return list.size
}

func (list *DoublyLinkedList) InsertAtLast(Data interface{}) {
	node := &DoubleNode{
		Data: Data,
		next: nil,
		prev: nil,
	}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		lastElement := list.Tail
		list.Tail.next = node
		node.prev = lastElement
		list.Tail = node
	}
	list.size++
}

func (list *DoublyLinkedList) InsertAtStart(Data interface{}) {
	node := &DoubleNode{
		Data: Data,
	}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.next = list.Head
		list.Head = node
	}
	list.size++
}

func (list *DoublyLinkedList) DeleteElement(Data interface{}) {
	if list.Head == nil {
		return
	}
	current := list.Head
	if current.Data == Data {
		//found at head
		current.prev = nil
		list.Head = current.next
		list.Head.prev = nil
		list.size--
		return
	}
	for current != nil {
		if current.Data == Data {
			//found element
			prev := current.prev
			next := current.next
			prev.next = next
			if next != nil {
				next.prev = prev
			} else {
				list.Tail = prev
			}
			list.size--
		}
		current = current.next
	}
}

func (list *DoublyLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.next
	}
}

func (list *DoublyLinkedList) InsertAfter(nodeElement interface{}, data interface{}) {
	current := list.Head
	node := &DoubleNode{
		Data: data,
	}
	for current != nil {
		if current.Data == nodeElement {
			next := current.next
			node.next = next
			node.prev = current
			if node.next == nil {
				list.Tail = node
			}
			list.size++
			return
		}
		current = current.next
	}
	fmt.Print("element not found , not inserting")
}
