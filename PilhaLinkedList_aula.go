package main

import (
	"fmt"
	"errors"
)

type Stack interface{
	Push(e int)
	Pop() (int, error)
	Peek() (int, error)
	size() int
}

type ArrayStack struct {
	v []int
	inserted int
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

func(LinkedListStack *S) Push(e int){
	newNode = &Node1P{v:e, next:S.Peek}
	S.Peek = newNode
	S.inserted++
}

func(S *ArrayStack) Pop()(int, error){
	if S.inserted == 0{return -1, error.New("Index invalido -- Erro no Pop")}
	S.inserted--
	aux := S.Peek.v
	S.Peek = S.Peek.next
	return aux, nil
}

func(S *ArrayStack) Peek()(int, error){
	if S.inserted == 0{return -1, error.New("Index invalido -- Erro no Peek")}
	return S.v[S.inserted[-1], nil]
}

