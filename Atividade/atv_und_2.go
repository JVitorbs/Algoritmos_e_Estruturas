package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Função para criar vetor de inteiros
func V(n int, sorted bool) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000000)
	}
	if sorted {
		sort.Ints(arr)
	}
	return arr
}

// Selection Sort com medição de tempo
func selectionSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	duration := time.Since(start)
	return arr, duration
}

// Bubble Sort com medição de tempo
func bubbleSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	duration := time.Since(start)
	return arr, duration
}

// Insertion Sort com medição de tempo
func insertionSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	duration := time.Since(start)
	return arr, duration
}

// Merge Sort com medição de tempo
func mergeSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	if len(arr) <= 1 {
		return arr, time.Since(start)
	}
	mid := len(arr) / 2
	left, _ := mergeSort(arr[:mid])
	right, _ := mergeSort(arr[mid:])
	result := merge(left, right)
	duration := time.Since(start)
	return result, duration
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Quick Sort (sem randomização) com medição de tempo
func quickSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	if len(arr) < 2 {
		return arr, time.Since(start)
	}
	stack := [][]int{{0, len(arr) - 1}}
	for len(stack) > 0 {
		low, high := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		
		if low < high {
			pivot := arr[high]
			i := low
			for j := low; j < high; j++ {
				if arr[j] < pivot {
					arr[i], arr[j] = arr[j], arr[i]
					i++
				}
			}
			arr[i], arr[high] = arr[high], arr[i]
			
			stack = append(stack, []int{low, i - 1})
			stack = append(stack, []int{i + 1, high})
		}
	}
	duration := time.Since(start)
	return arr, duration
}

// Quick Sort (com randomização) com medição de tempo
func quickSortRandom(arr []int) ([]int, time.Duration) {
	start := time.Now()
	if len(arr) < 2 {
		return arr, time.Since(start)
	}
	stack := [][]int{{0, len(arr) - 1}}
	for len(stack) > 0 {
		low, high := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		
		if low < high {
			pivotIndex := rand.Intn(high-low+1) + low
			arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
			pivot := arr[high]
			i := low
			for j := low; j < high; j++ {
				if arr[j] < pivot {
					arr[i], arr[j] = arr[j], arr[i]
					i++
				}
			}
			arr[i], arr[high] = arr[high], arr[i]
			
			stack = append(stack, []int{low, i - 1})
			stack = append(stack, []int{i + 1, high})
		}
	}
	duration := time.Since(start)
	return arr, duration
}

// Counting Sort com medição de tempo
func countingSort(arr []int) ([]int, time.Duration) {
	start := time.Now()
	if len(arr) == 0 {
		return arr, time.Since(start)
	}
	
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}
	
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	
	sorted := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		pos := count[num] - 1
		sorted[pos] = num
		count[num]--
	}
	duration := time.Since(start)
	return sorted, duration
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Criar vetor com 10^6 elementos
	const n = 1000000
	fmt.Println("Criando vetores...")
	unsortedArr := V(n, false)
	sortedArr := make([]int, n)
	copy(sortedArr, unsortedArr)
	sort.Ints(sortedArr)
	
	// Lista de algoritmos para testar
	algorithms := []struct {
		name string
		fn   func([]int) ([]int, time.Duration)
	}{
		{"Selection Sort", selectionSort},
		{"Bubble Sort", bubbleSort},
		{"Insertion Sort", insertionSort},
		{"Merge Sort", mergeSort},
		{"Quick Sort (sem random)", quickSort},
		{"Quick Sort (com random)", quickSortRandom},
		{"Counting Sort", countingSort},
	}
	
	// Testar com vetor desordenado
	fmt.Println("\nTestando com vetor DESORDENADO:")
	for _, algo := range algorithms {
		_, duration := algo.fn(unsortedArr)
		fmt.Printf("%-22s → %v\n", algo.name, duration)
	}
	
	// Testar com vetor ordenado
	fmt.Println("\nTestando com vetor ORDENADO:")
	for _, algo := range algorithms {
		_, duration := algo.fn(sortedArr)
		fmt.Printf("%-22s → %v\n", algo.name, duration)
	}
}