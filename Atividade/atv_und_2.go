package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// Selection Sort
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// Bubble Sort
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
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
	return arr
}

// Insertion Sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		L := make([]int, mid)
		R := make([]int, len(arr)-mid)

		copy(L, arr[:mid])
		copy(R, arr[mid:])

		mergeSort(L)
		mergeSort(R)

		i, j, k := 0, 0, 0

		for i < len(L) && j < len(R) {
			if L[i] < R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}
			k++
		}

		for i < len(L) {
			arr[k] = L[i]
			i++
			k++
		}

		for j < len(R) {
			arr[k] = R[j]
			j++
			k++
		}
	}
	return arr
}

// Quick Sort (sem randomização)
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, right []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

// Quick Sort (com randomização)
func randomizedQuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	randIdx := rand.Intn(len(arr))
	pivot := arr[randIdx]
	arr[randIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[randIdx]

	var left, right []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = randomizedQuickSort(left)
	right = randomizedQuickSort(right)

	return append(append(left, pivot), right...)
}

// Counting Sort (para inteiros não negativos)
func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	maxVal := arr[0]
	minVal := arr[0]
	for _, num := range arr {
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	count := make([]int, maxVal-minVal+1)
	output := make([]int, len(arr))

	for _, num := range arr {
		count[num-minVal]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]-minVal]-1] = arr[i]
		count[arr[i]-minVal]--
	}

	return output
}

// Função para gerar vetores
func generateArrays(size int, ordered bool, descending bool) []int {
	arr := make([]int, size)

	if ordered {
		for i := 0; i < size; i++ {
			if descending {
				arr[i] = size - i
			} else {
				arr[i] = i + 1
			}
		}
	} else {
		for i := 0; i < size; i++ {
			arr[i] = i + 1
		}
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	return arr
}

// Função para testar os algoritmos
func testAlgorithms(algorithms []func([]int) []int, sizes []int, cases map[string]struct {
	ordered    bool
	descending bool
}) map[string]map[string][]float64 {
	results := make(map[string]map[string][]float64)
	for _, alg := range algorithms {
		results[getSimpleFunctionName(alg)] = make(map[string][]float64)
		for caseName := range cases {
			results[getSimpleFunctionName(alg)][caseName] = make([]float64, len(sizes))
		}
	}

	for i, size := range sizes {
		fmt.Printf("\nTestando com tamanho: %d\n", size)
		for caseName, params := range cases {
			arr := generateArrays(size, params.ordered, params.descending)

			for _, algorithm := range algorithms {
				testArr := make([]int, len(arr))
				copy(testArr, arr)

				algName := getSimpleFunctionName(algorithm)

				startTime := time.Now()
				algorithm(testArr)
				elapsedTime := time.Since(startTime).Seconds()

				results[algName][caseName][i] = elapsedTime
				fmt.Printf("Algoritmo: %-18s | Caso: %-20s | Tamanho: %-8d | Tempo: %12.6f segundos\n",
					algName, caseName, size, elapsedTime)
			}
		}
	}
	return results
}

// Função corrigida para obter o nome da função
func getFunctionName(f func([]int) []int) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Função para extrair apenas o nome simples da função (sem o caminho do pacote)
func getSimpleFunctionName(f func([]int) []int) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	// Extrai apenas a parte após o último ponto
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			return fullName[i+1:]
		}
	}
	return fullName
}

func main() {
	rand.Seed(time.Now().UnixNano())

	algorithms := []func([]int) []int{
		selectionSort,
		bubbleSort,
		insertionSort,
		mergeSort,
		quickSort,
		randomizedQuickSort,
		countingSort,
	}

	sizes := []int{int(math.Pow10(5)), int(math.Pow10(6))}

	cases := map[string]struct {
		ordered    bool
		descending bool
	}{
		"Ordenado Crescente":   {true, false},
		"Ordenado Decrescente": {true, true},
		"Desordenado":          {false, false},
	}

	results := testAlgorithms(algorithms, sizes, cases)

	// Exibir resultados em tabela
	fmt.Println("\nResultados consolidados:")
	for alg, algCases := range results {
		fmt.Printf("\nAlgoritmo: %s\n", alg)
		fmt.Println("Caso\t\tTamanho 1e3\tTamanho 1e4\tTamanho 1e5")
		for caseName, times := range algCases {
			fmt.Printf("%-15s\t%.6f\t%.6f\t%.6f\n", caseName, times[0], times[1], times[2])
		}
	}
}
