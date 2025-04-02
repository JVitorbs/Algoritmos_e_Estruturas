package main

import "fmt"

// Definição da estrutura da Pilha
type Stack struct {
    items []int
}

// Método para adicionar um elemento na pilha (push)
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

// Método para remover e retornar o topo da pilha (pop)
func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        fmt.Println("Pilha vazia!")
        return -1
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top
}

func main() {
    stack := Stack{}
    stack.Push(10)
    stack.Push(20)
    fmt.Println(stack.Pop()) // Saída: 20
    fmt.Println(stack.Pop()) // Saída: 10
}
