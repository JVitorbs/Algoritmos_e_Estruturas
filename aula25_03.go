// Codigo da aula do dia 25.03.2025
package main

import (
	"fmt"
	//"errors"
)

type List interface {
	Size() int
	Get(index int) int
	Add(e int)
	AddOnIndex(e int, index int)
	Remove(index int)
}

type ArrayList struct {
	v        []int
	inserted int
}

func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

func (l *ArrayList) Get(index int) int { //O(4) = O(1), constante
	if index >= 0 && index < l.inserted {
		return l.v[index]
	} else {
		return -1 //error
	}

}

func (l *ArrayList) Size() int {
	return l.inserted
}

func (l *ArrayList) doubleV() {
	newV := make([]int, 2*l.inserted)
	for i := 0; i < l.inserted; i++ {
		newV[i] = l.v[i]
	}
	l.v = newV
}

func (l *ArrayList) Add(e int) {
	if l.inserted == len(l.v) {
		l.doubleV()
	}
	l.v[l.inserted] = e
	l.inserted++
}

func (l *ArrayList) AddOnIndex(e int, index int) {
	if index < 0 || index > l.inserted {
		return
	}

	if l.inserted == len(l.v) {
		l.doubleV()
	}

	for i := l.inserted; i > index; i-- {
		l.v[i] = l.v[i-1]
	}
	l.v[index] = e
	l.inserted++

}

func (l *ArrayList) Remove(index int) {
	if index >= 0 && index < l.inserted {
		for i := index; ; i++ {

		}
	}
}

func main() {
	l := &ArrayList{}
	l.Init(10)
	for i := 1; i <= 50; i++ {
		l.Add(i)
	}
	l.AddOnIndex(0, 0)
	fmt.Println(l.Get(49))

}
