package main

import (
	"fmt"
)

// Estrutura do nó da BST com altura e fator de balanço
type BSTNode struct {
	Value         int
	Left, Right   *BSTNode
	Height        int
	BalanceFactor int
}

// Função auxiliar para obter altura de um nó
func height(node *BSTNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// Atualiza altura e fator de balanço
func update(node *BSTNode) {
	if node == nil {
		return
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	node.Height = 1 + max(leftHeight, rightHeight)
	node.BalanceFactor = leftHeight - rightHeight
}

// Inserir um valor na árvore
func Add(node *BSTNode, value int) *BSTNode {
	if node == nil {
		return &BSTNode{Value: value, Height: 1}
	}
	if value < node.Value {
		node.Left = Add(node.Left, value)
	} else if value > node.Value {
		node.Right = Add(node.Right, value)
	} else {
		return node // ignora duplicados
	}
	update(node)
	return node
}

// Buscar um valor na árvore
func Find(node *BSTNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return Find(node.Left, value)
	}
	return Find(node.Right, value)
}

// Encontrar o menor valor na subárvore (usado na remoção)
func findMin(node *BSTNode) *BSTNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Remover um valor da árvore
func Remove(node *BSTNode, value int) *BSTNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = Remove(node.Left, value)
	} else if value > node.Value {
		node.Right = Remove(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minBSTNode := findMin(node.Right)
		node.Value = minBSTNode.Value
		node.Right = Remove(node.Right, minBSTNode.Value)
	}
	update(node)
	return node
}

// Percurso em ordem (in-order)
func InOrder(node *BSTNode) {
	if node != nil {
		InOrder(node.Left)
		fmt.Printf("%d (H:%d, B:%d) ", node.Value, node.Height, node.BalanceFactor)
		InOrder(node.Right)
	}
}

// Função utilitária
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Função principal
func main() {
	var root *BSTNode
	values := []int{40, 20, 10, 30, 60, 50, 70}

	for _, v := range values {
		root = Add(root, v)
	}

	fmt.Print("InOrder com altura e fator de balanço: ")
	InOrder(root)
	fmt.Println()

	fmt.Println("Buscar 30:", Find(root, 30))   // true
	fmt.Println("Buscar 100:", Find(root, 100)) // false

	root = Remove(root, 20)
	fmt.Print("Após remover 20: ")
	InOrder(root)
	fmt.Println()
}
