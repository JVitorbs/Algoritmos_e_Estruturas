package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"time"
)

// 1. Algoritmos de Ordenação

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

	pivot := arr[len(arr)-1]
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

// 2. Funções Auxiliares

// getFunctionName obtém o nome da função
func getFunctionName(f func([]int) []int) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// generateArrays gera vetores de teste
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

// testAlgorithms testa todos os algoritmos
func testAlgorithms(algorithms []func([]int) []int, sizes []int, cases map[string]struct {
	ordered    bool
	descending bool
}) map[string]map[string][]float64 {
	
	results := make(map[string]map[string][]float64)
	for _, alg := range algorithms {
		algName := getFunctionName(alg)
		results[algName] = make(map[string][]float64)
		for caseName := range cases {
			results[algName][caseName] = make([]float64, len(sizes))
		}
	}

	for i, size := range sizes {
		fmt.Printf("\nTestando com tamanho: %d\n", size)
		for caseName, params := range cases {
			arr := generateArrays(size, params.ordered, params.descending)
			
			for _, algorithm := range algorithms {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				
				startTime := time.Now()
				algorithm(testArr)
				elapsedTime := time.Since(startTime).Seconds()
				
				algName := getFunctionName(algorithm)
				results[algName][caseName][i] = elapsedTime
				fmt.Printf("%-20s - %-15s - Tamanho %d: %.6f segundos\n", 
					algName, caseName, size, elapsedTime)
			}
		}
	}
	return results
}

// 3. Relatório e Visualização

// printResults imprime os resultados em formato de tabela
func printResults(results map[string]map[string][]float64, sizes []int) {
	fmt.Println("\nRESULTADOS:")
	fmt.Println("==============================================")
	
	for caseName := range results[getFunctionName(selectionSort)] {
		fmt.Printf("\nCaso: %s\n", caseName)
		fmt.Printf("%-30s", "Algoritmo")
		for _, size := range sizes {
			fmt.Printf(" | %8d", size)
		}
		fmt.Println("\n" + repeatChar("-", 30 + len(sizes)*10))
		
		for algName, algCases := range results {
			fmt.Printf("%-30s", algName)
			for _, t := range algCases[caseName] {
				fmt.Printf(" | %8.4f", t)
			}
			fmt.Println()
		}
	}
}

// repeatChar auxiliar para criar linhas
func repeatChar(char string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}

// generateReport gera um relatório em arquivo de texto
func generateReport(results map[string]map[string][]float64, sizes []int) {
	file, err := os.Create("relatorio.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "RELATÓRIO DE DESEMPENHO DE ALGORITMOS DE ORDENAÇÃO")
	fmt.Fprintln(file, "=================================================")
	
	// Complexidade teórica
	fmt.Fprintln(file, "\nCOMPLEXIDADE TEÓRICA:")
	fmt.Fprintln(file, "Selection Sort: O(n²) em todos os casos")
	fmt.Fprintln(file, "Bubble Sort: O(n) melhor caso, O(n²) pior caso")
	fmt.Fprintln(file, "Insertion Sort: O(n) melhor caso, O(n²) pior caso")
	fmt.Fprintln(file, "Merge Sort: O(n log n) em todos os casos")
	fmt.Fprintln(file, "Quick Sort: O(n log n) médio, O(n²) pior caso")
	fmt.Fprintln(file, "Quick Sort Randomizado: O(n log n) em todos os casos")
	fmt.Fprintln(file, "Counting Sort: O(n + k) onde k é o intervalo de valores")
	
	// Resultados
	for caseName := range results[getFunctionName(selectionSort)] {
		fmt.Fprintf(file, "\nCASO: %s\n", caseName)
		fmt.Fprintf(file, "%-30s", "Algoritmo")
		for _, size := range sizes {
			fmt.Fprintf(file, " | %8d", size)
		}
		fmt.Fprintln(file, "\n" + repeatChar("-", 30 + len(sizes)*10))
		
		for algName, algCases := range results {
			fmt.Fprintf(file, "%-30s", algName)
			for _, t := range algCases[caseName] {
				fmt.Fprintf(file, " | %8.4f", t)
			}
			fmt.Fprintln(file)
		}
	}
	
	// Análise
	fmt.Fprintln(file, "\nANÁLISE:")
	fmt.Fprintln(file, "1. Algoritmos O(n²) são adequados apenas para pequenos conjuntos")
	fmt.Fprintln(file, "2. Merge Sort mostra desempenho consistente em todos os casos")
	fmt.Fprintln(file, "3. Quick Sort padrão tem pior desempenho em vetores ordenados")
	fmt.Fprintln(file, "4. Quick Sort randomizado evita o pior caso")
	fmt.Fprintln(file, "5. Counting Sort é o mais rápido quando aplicável")
	
	fmt.Println("\nRelatório gerado como relatorio.txt")
}

// 4. Função principal
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

	sizes := []int{int(math.Pow10(3)), int(math.Pow10(4)), int(math.Pow10(5))}

	cases := map[string]struct {
		ordered    bool
		descending bool
	}{
		"Ordenado Crescente":  {true, false},
		"Ordenado Decrescente": {true, true},
		"Desordenado":         {false, false},
	}

	results := testAlgorithms(algorithms, sizes, cases)

	// Exibir resultados
	printResults(results, sizes)

	// Gerar relatório
	generateReport(results, sizes)
}