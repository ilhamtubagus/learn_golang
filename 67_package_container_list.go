package main

import (
	"container/list"
	"fmt"
)

// Implementation of double linked list in golang

func main() {
	var data *list.List = list.New()

	// add elements to the end of the list
	data.PushBack("Ilham")
	data.PushBack("Tubagus")
	data.PushBack("Arfian")

	var head *list.Element = data.Front()
	var tail *list.Element = data.Back()
	fmt.Println("head", head.Value)
	fmt.Println("tail", tail.Value)

	fmt.Println("tail", head.Next().Next().Value)

	// looping from head to tail
	for e := data.Front(); e != nil; e = e.Next() {
		println(e.Value.(string))
	}

	// looping from tail to head
	for e := data.Back(); e != nil; e = e.Prev() {
		println(e.Value.(string))
	}
}
