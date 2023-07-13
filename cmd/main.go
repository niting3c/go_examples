package main

import (
	"fmt"

	"github.com/niting3c/go_examples/linkedlists/singly"
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
}
