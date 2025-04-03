package main

import "fmt"

// Estrutura do nó da lista
type Node struct {
	prev  *Node
	value int
	next  *Node
}

// Estrutura da lista duplamente encadeada
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Método para adicionar um nó ao final da lista
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &Node{value: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

// Método para remover um nó pelo valor
func (dll *DoublyLinkedList) Remove(value int) {
	current := dll.head

	for current != nil {
		if current.value == value {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dll.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dll.tail = current.prev
			}
			return
		}
		current = current.next
	}
}

// Método para exibir os elementos da lista
func (dll *DoublyLinkedList) Display() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	dll := &DoublyLinkedList{}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	fmt.Println("Lista após inserções:")
	dll.Display()

	dll.Remove(20)
	fmt.Println("Lista após remoção de 20:")
	dll.Display()
}
