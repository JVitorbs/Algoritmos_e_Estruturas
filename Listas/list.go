package main

import (
	"fmt"
)

// Interface List
type List interface {
	Size() int
	Get(index int) int
	Add(e int)
	AddOnIndex(e int, index int)
	Remove(index int)
}

// Estrutura ArrayList
type ArrayList struct {
	v        []int
	inserted int
}

// Inicializa a lista com um tamanho fixo
func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

// Retorna um elemento pelo índice
func (l *ArrayList) Get(index int) int {
	if index >= 0 && index < l.inserted {
		return l.v[index]
	}
	return -1 // erro (ideal seria retornar um erro)
}

// Retorna o tamanho da lista
func (l *ArrayList) Size() int {
	return l.inserted
}

// Dobra o tamanho do slice quando necessário
func (l *ArrayList) doubleV() {
	newV := make([]int, 2*len(l.v))
	copy(newV, l.v)
	l.v = newV
}

// Adiciona um elemento no final da lista
func (l *ArrayList) Add(e int) {
	if l.inserted == len(l.v) {
		l.doubleV()
	}
	l.v[l.inserted] = e
	l.inserted++
}

// Adiciona um elemento em um índice específico
func (l *ArrayList) AddOnIndex(e int, index int) {
	if index < 0 || index > l.inserted {
		fmt.Println("Índice inválido")
		return
	}

	if l.inserted == len(l.v) {
		l.doubleV()
	}

	// Desloca os elementos para abrir espaço
	for i := l.inserted; i > index; i-- {
		l.v[i] = l.v[i-1]
	}
	l.v[index] = e
	l.inserted++
}

// Remove um elemento de um índice específico
func (l *ArrayList) Remove(index int) {
	if index < 0 || index >= l.inserted {
		fmt.Println("Índice inválido")
		return
	}

	// Desloca os elementos para fechar o espaço
	for i := index; i < l.inserted-1; i++ {
		l.v[i] = l.v[i+1]
	}
	l.inserted--
}

// Função principal para testar
func main() {
	l := &ArrayList{}
	l.Init(5)

	// Testando Add()
	for i := 1; i <= 5; i++ {
		l.Add(i)
	}
	fmt.Println("Lista após adicionar elementos:", l.v[:l.inserted])

	// Testando AddOnIndex()
	l.AddOnIndex(99, 2)
	fmt.Println("Lista após adicionar 99 no índice 2:", l.v[:l.inserted])

	// Testando Remove()
	l.Remove(3)
	fmt.Println("Lista após remover o índice 3:", l.v[:l.inserted])
}
