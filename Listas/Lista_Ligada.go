package main

import "fmt"

// Definição do nó da lista
type Node struct {
	data int
	next *Node
}

// Definição da lista ligada
type LinkedList struct {
	head *Node
}

// Método para adicionar um novo nó no final
func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Método para imprimir a lista
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Print() // Saída: 10 -> 20 -> 30 -> nil
}
