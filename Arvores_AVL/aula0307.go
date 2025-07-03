package main

import (
	"fmt"
)

// Definição do nó da árvore
type Node struct {
	Value int
	Left  *Node
	Right *Node
	height int
	BalanFact int
}

// Inserir um valor na árvore
func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value > root.Value {
		root.Right = Insert(root.Right, value)
	}
	return root
}

// Buscar um valor na árvore
func Search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if value == root.Value {
		return true
	} else if value < root.Value {
		return Search(root.Left, value)
	}
	return Search(root.Right, value)
}

// Encontrar o menor valor (usado na remoção)
func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Remover um valor da árvore
func Remove(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if value < root.Value {
		root.Left = Remove(root.Left, value)
	} else if value > root.Value {
		root.Right = Remove(root.Right, value)
	} else {
		// Caso com 1 ou nenhum filho
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// Caso com dois filhos
		minNode := findMin(root.Right)
		root.Value = minNode.Value
		root.Right = Remove(root.Right, minNode.Value)
	}
	return root
}

// Percurso em ordem (in-order)
func InOrder(root *Node) {
	if root != nil {
		InOrder(root.Left)
		fmt.Print(root.Value, " ")
		InOrder(root.Right)
	}
}



// Função principal
func main() {
	var root *Node

	// Inserções
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		root = Insert(root, v)
	}

	fmt.Print("In-Order: ")
	InOrder(root)
	fmt.Println()

	// Busca
	fmt.Println("Buscar 40:", Search(root, 40)) // true
	fmt.Println("Buscar 100:", Search(root, 100)) // false

	// Remoção
	root = Remove(root, 70)
	fmt.Print("Após remover 70: ")
	InOrder(root)
	fmt.Println()
}