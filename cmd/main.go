package main

import (
	"fmt"

	doubly "github.com/niting3c/go_examples/linkedlists/doubly"
	singly "github.com/niting3c/go_examples/linkedlists/singly"
)

func main() {
	list := &singly.SinglyLinkedList{}
	list.InsertAtLast(1)
	list.InsertAtLast(2)
	list.InsertAtLast(3)

	fmt.Println("Head:", list.Head.Data)
	fmt.Println("Tail:", list.Tail.Data)
	fmt.Println("Size:", list.Size())

	list.InsertAtStart(0)
	list.Print()

	list.DeleteElement(2)
	list.Print()
	fmt.Println("-----------------")
	fmt.Println("Printing for doubly linked List")

	doubleList := &doubly.DoublyLinkedList{}
	doubleList.InsertAtLast(1)
	doubleList.InsertAtLast(2)
	doubleList.InsertAtLast(3)

	fmt.Println("Head:", doubleList.Head.Data)
	fmt.Println("Tail:", doubleList.Tail.Data)
	fmt.Println("Size:", doubleList.Size())

	doubleList.InsertAtStart(0)
	doubleList.Print()

	doubleList.DeleteElement(2)
	doubleList.Print()
}
