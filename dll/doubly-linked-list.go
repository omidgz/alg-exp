package main

import "fmt"

type Data struct {
	ID    int
	Value string
}

type Node struct {
	value      Data
	prev, next *Node
}

type DLL struct {
	size        int
	length      int
	first, last *Node
}

var ErrListFull = fmt.Errorf("List full")

func (this *DLL) grow() {
	this.length++
}

func (this *DLL) shrink() {
	this.length--
}

func (this *DLL) PutLast(value Data) error {
	if found := this.Find(value.ID); found != nil {
		found.value = value
		return nil
	}

	if this.length == this.size {
		return ErrListFull
	}

	newNode := &Node{
		value: value,
	}
	if this.last == nil {
		this.first = newNode
		this.last = newNode
		this.grow()
		return nil
	}

	newNode.prev = this.last
	this.last.next = newNode
	this.last = newNode
	this.grow()
	return nil
}

func (this *DLL) PutFirst(value Data) error {
	if found := this.Find(value.ID); found != nil {
		found.value = value
		return nil
	}

	if this.length == this.size {
		return ErrListFull
	}

	newNode := &Node{
		value: value,
	}

	if this.first == nil {
		this.first = newNode
		this.last = newNode
		this.grow()
		return nil
	}

	newNode.next = this.first
	this.first.prev = newNode
	this.first = newNode
	this.grow()
	return nil
}

func (this *DLL) Find(id int) *Node {
	curr := this.first
	for curr != nil {
		if curr.value.ID == id {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (this *DLL) PrintLTR() {
	curr := this.first
	for curr != nil {
		fmt.Printf("%v ", curr.value)
		curr = curr.next
	}
	fmt.Println(this.length)
}

func (this *DLL) PrintRTL() {
	curr := this.last
	for curr != nil {
		fmt.Printf("%v ", curr.value)
		curr = curr.prev
	}
	fmt.Println(this.length)
}

func main() {
	dll := DLL{
		size: 5,
	}

	fmt.Println(dll.PutLast(Data{ID: 1, Value: "data1"}))
	dll.PrintLTR()
	fmt.Println(dll.PutLast(Data{ID: 2, Value: "data2"}))
	dll.PrintLTR()
	fmt.Println(dll.PutLast(Data{ID: 3, Value: "data3"}))
	dll.PrintLTR()
	fmt.Println(dll.PutFirst(Data{ID: 4, Value: "data4"}))
	dll.PrintRTL()
	fmt.Println(dll.PutFirst(Data{ID: 3, Value: "data3-2"}))
	dll.PrintRTL()
	dll.PrintLTR()
	fmt.Println(dll.PutLast(Data{ID: 5, Value: "data5"}))
	dll.PrintLTR()
	fmt.Println(dll.PutLast(Data{ID: 6, Value: "data6"}))
	dll.PrintLTR()
}
