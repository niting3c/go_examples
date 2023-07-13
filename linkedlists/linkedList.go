package singly

type LinkedList interface {
	InsertAtLast(Data interface{})
	Size() int64
	InsertAtStart(Data interface{})
	DeleteElement(Data interface{})
	Print()
	InsertAfter(nodeElement interface{}, Data interface{})
}
