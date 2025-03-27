package main

import (
	"fmt"
	"errors"
)

type List interface {
	Size() int
	Get(index int) (int, error)
	Add(e int) error
	AddOnIndex(e int, index int) error
	Remove(index int) error
}

type ArrayList struct {
	v []int
	inserted int
}

func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

func (l *ArrayList) Get(index int) (int,error) { //O(1),Ômega(1), então Theta(1)
	if index >= 0 && index < l.inserted {	
		return l.v[index], nil		
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	
}

func (l *ArrayList) Size() int { //Theta(1)
	return l.inserted
}

func (l *ArrayList) doubleV() { //Theta(n)
	newV := make([]int,2*l.inserted)
	for i := 0; i < l.inserted; i++{
		newV[i] = l.v[i]
	}
	l.v = newV
}

func (l *ArrayList) Add(e int) { //Ômega(1), O(n)
	if l.inserted == len(l.v){
		l.doubleV()
	}
	l.v[l.inserted] = e
	l.inserted++
}

func (l *ArrayList) AddOnIndex(e int, index int) error { //Ômega(1), O(n)
	if index < 0 || index > l.inserted {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	if l.inserted == len(l.v){
		l.doubleV()
	}
	for i := l.inserted; i>index ;i-- {
		l.v[i] = l.v[i-1]
	}
	l.v[index] = e
	l.inserted++
	return nil
}

func (l *ArrayList) Remove(index int) error {	//Ômega(1), O(n)
	if index >=0 && index < l.inserted {
		for i:=index; i < l.inserted - 1; i++ {
			l.v[i] = l.v[i+1]
		}
		l.inserted--
		return nil
	} else{
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

func main(){
	l := &ArrayList{}
	l.Init(10)
	for i:=1; i <= 50; i++{
		l.Add(i)
	}
	erro := l.AddOnIndex(-1,-1)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
	val,erro := l.Get(49)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
	fmt.Println(val)
	erro = l.Remove(0)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
	val,erro = l.Get(49)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
}
