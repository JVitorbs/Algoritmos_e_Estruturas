package main

import "fmt"

// Estrutura da Fila
type Queue struct {
	items []int
}

// Método para adicionar na fila (enqueue)
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Método para remover da fila (dequeue)
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Fila vazia!")
		return -1
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	fmt.Println(queue.Dequeue()) // Saída: 10
	fmt.Println(queue.Dequeue()) // Saída: 20
}
