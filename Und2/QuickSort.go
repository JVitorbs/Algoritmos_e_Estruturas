package main

import "fmt"

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// Encontra a posição correta do pivô
		pivotIndex := partition(arr, low, high)

		// Ordena recursivamente os elementos antes e depois do pivô
		quickSortRecursive(arr, low, pivotIndex-1)
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Escolhe o último elemento como pivô
	i := low           // Índice do menor elemento

	for j := low; j < high; j++ {
		// Se o elemento atual é menor ou igual ao pivô
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i] // Troca os elementos
			i++                             // Incrementa o índice do menor elemento
		}
	}
	// Coloca o pivô na posição correta
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main() {
	arr := []int{9, 4, 3, 6, 3, 2, 8, 7, 1, 5}
	fmt.Println("Array original:", arr)

	quickSortRecursive(arr, 0, len(arr)-1)
	fmt.Println("Array ordenado:", arr)
}