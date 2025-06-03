package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*Função para gerar vetores ordenados ou não
  Cria um slice de inteiros com valores aleatórios
  Pode gerar já ordenado ou não, conforme parâmetro 
*/
func V(n int, sorted bool) []int {
	arr := make([]int, n) // Cria um slice com capacidade n
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000000) // Preenche com valores aleatórios até 1.000.000
	}
	if sorted {
		sort.Ints(arr) // Ordena se solicitado
	}
	return arr
}

/*selectionSort
  Complexidade: O(n²) em todos os casos
  Funcionamento: Seleciona repetidamente o menor elemento 
*/
func selectionSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now() // Marca o início da execução
	n := len(arr)
	
	for i := 0; i < n-1; i++ {
		minIndex := i // Assume que o atual é o menor
		
		// Procura o menor elemento no restante do array
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j // Encontrou um menor
			}
		}
		
		// Troca o menor elemento encontrado com a posição atual
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	
	return arr, time.Since(comeco) // Retorna array ordenado e tempo gasto
}

/* bubbleSort
   Complexidade: O(n²) no pior caso, O(n) no melhor caso
   Funcionamento: Troca repetidamente elementos adjacentes desordenados
*/
func bubbleSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	n := len(arr)
	
	for i := 0; i < n-1; i++ {
		swapped := false // Flag para otimização
		
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Troca elementos adjacentes
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		
		// Se não houve trocas, o array já está ordenado
		if !swapped {
			break
		}
	}
	
	return arr, time.Since(comeco)
}

/* insertionSort
   Complexidade: O(n²) no pior caso, O(n) no melhor caso
   Funcionamento: Constrói a sequência ordenada um item de cada vez
*/
func insertionSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	n := len(arr)
	
	for i := 1; i < n; i++ {
		key := arr[i] // Elemento atual a ser inserido
		j := i - 1    // Índice do elemento anterior
		
		// Move elementos maiores que key para frente
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key // Insere key na posição correta
	}
	
	return arr, time.Since(comeco)
}

/* mergeSort
   Complexidade: O(n log n) em todos os casos
   Funcionamento: Dividir e conquistar com recursão
*/
func mergeSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	
	// Caso base: arrays com 0 ou 1 elemento já estão ordenados
	if len(arr) <= 1 {
		return arr, time.Since(comeco)
	}
	
	// Divide o array em duas partes
	mid := len(arr) / 2
	
	// Ordena recursivamente cada metade
	left, _ := mergeSort(arr[:mid])
	right, _ := mergeSort(arr[mid:])
	
	// Combina as duas metades ordenadas
	result := merge(left, right)
	
	return result, time.Since(comeco)
}

// Função auxiliar para combinar dois subarrays ordenados
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0 // Índices para percorrer left e right
	
	// Combina ordenadamente enquanto houver elementos em ambos
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	// Adiciona os elementos restantes (se houver)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

/* quickSort
   Complexidade: O(n²) no pior caso, O(n log n) em média
   Funcionamento: Particionamento em torno de um pivô
*/
func quickSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	
	// Caso base: arrays pequenos já estão ordenados
	if len(arr) < 2 {
		return arr, time.Since(comeco)
	}
	
	// Implementação iterativa usando stack para evitar recursão
	stack := [][]int{{0, len(arr) - 1}} // Stack com intervalos a ordenar
	
	for len(stack) > 0 {
		// Pega o próximo intervalo da stack
		low, high := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		
		if low < high {
			// Particiona o array
			pivot := arr[high] // Escolhe o último elemento como pivô
			i := low           // Índice do menor elemento
			
			for j := low; j < high; j++ {
				if arr[j] < pivot {
					arr[i], arr[j] = arr[j], arr[i]
					i++
				}
			}
			
			// Coloca o pivô na posição correta
			arr[i], arr[high] = arr[high], arr[i]
			
			// Adiciona os subarrays à stack para ordenação
			stack = append(stack, []int{low, i - 1})
			stack = append(stack, []int{i + 1, high})
		}
	}
	
	return arr, time.Since(comeco)
}

/* quickSort com pivô aleatorio
   Complexidade: O(n log n) na maioria dos casos
   Funcionamento: Melhora a versão clássica com pivô aleatório
*/
func quickSortRandom(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	
	if len(arr) < 2 {
		return arr, time.Since(comeco)
	}
	
	stack := [][]int{{0, len(arr) - 1}}
	
	for len(stack) > 0 {
		low, high := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		
		if low < high {
			// Seleciona um pivô aleatório e o move para o final
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
	
	return arr, time.Since(comeco)
}

/* countingSort
   Complexidade: O(n + k) onde k é o valor máximo
   Funcionamento: Não baseado em comparações, ideal para números em faixa limitada
*/
func countingSort(arr []int) ([]int, time.Duration) {
	comeco := time.Now()
	
	if len(arr) == 0 {
		return arr, time.Since(comeco)
	}
	
	// Encontra o valor máximo no array
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Cria array de contagem
	count := make([]int, max+1)
	
	// Conta a ocorrência de cada número
	for _, v := range arr {
		count[v]++
	}
	
	// Calcula as posições finais
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	
	// Constrói o array ordenado
	sorted := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		pos := count[num] - 1
		sorted[pos] = num
		count[num]--
	}
	
	return sorted, time.Since(comeco)
}

// Função main
func main() {
	// Configura a semente para geração de números aleatórios
	rand.Seed(time.Now().UnixNano())
	
	// Tamanho do vetor de teste (aumente para testes mais significativos)
    const n = 1000000
	fmt.Printf("\nCriando vetores com %d elementos...\n", n)
	
	// Gera vetor desordenado e uma cópia ordenada
	unsortedArr := V(n, false)
	sortedArr := make([]int, n)
	copy(sortedArr, unsortedArr)
	sort.Ints(sortedArr)
	
	// Lista de algoritmos para benchmark
	algorithms := []struct {
		name string               // Nome do algoritmo
		fn   func([]int) ([]int, time.Duration) // Função de ordenação
	}{
		{"Selection Sort", selectionSort},
		{"Bubble Sort", bubbleSort},
		{"Insertion Sort", insertionSort},
		{"Merge Sort", mergeSort},
		{"Quick Sort (sem random)", quickSort},
		{"Quick Sort (com random)", quickSortRandom},
		{"Counting Sort", countingSort},
	}
	
	/* EXECUÇÃO DOS TESTES COM VETOR DESORDENADO
	   Mostra o desempenho de cada algoritmo com dados aleatórios
	*/
	fmt.Println("\nTESTE COM VETOR DESORDENADO")
	for _, algo := range algorithms {
		// Executa a ordenação e mede o tempo
		_, duration := algo.fn(unsortedArr)
		
		// Formata a saída para melhor visualização
		fmt.Printf("%-22s → %v\n", algo.name, duration)
	}
	
	/* Execução dos testes com vetor ordenado
	   Mostra como cada algoritmo se comporta com dados já ordenados
    */
	fmt.Println("\nTESTE COM VETOR ORDENADO")
	for _, algo := range algorithms {
		_, duration := algo.fn(sortedArr)
		fmt.Printf("%-22s → %v\n", algo.name, duration)
	}
	
}