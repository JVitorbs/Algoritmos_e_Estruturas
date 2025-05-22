package main

import (
	"fmt"
	"errors"
)

type IStack interface{
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayStack struct {
	v []char
	size int
}

type Node1P struct{
	v int
	next * Node1P
}

type LinkedListStack struct{
	peek *Node1P
	inserted int
}

func (S *ArrayStack) Init(size int) {
	S.v = make([]int,7 size)
}

func (LinkedListStack *S) Push (value int){
	newNode = &Node1P{v:value}
}