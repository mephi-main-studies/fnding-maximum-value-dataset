package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	n := len(data)
	baseSize := n / CHUNKS
	remainder := n % CHUNKS

	var wg sync.WaitGroup
	maxima := make([]int, CHUNKS)
	valid := make([]bool, CHUNKS)

	start := 0
	for i := 0; i < CHUNKS; i++ {
		currentSize := baseSize
		if i < remainder {
			currentSize++
		}
		end := start + currentSize
		if currentSize > 0 { // запускаем горутину только для непустого куска
			s, e, idx := start, end, i
			wg.Add(1)
			go func() {
				defer wg.Done()
				localMax := data[s]
				for j := s + 1; j < e; j++ {
					if data[j] > localMax {
						localMax = data[j]
					}
				}
				maxima[idx] = localMax
				valid[idx] = true
			}()
		}
		start = end
	}

	wg.Wait()

	// Свести локальные максимумы к глобальному
	have := false
	var globalMax int
	for i := 0; i < CHUNKS; i++ {
		if valid[i] {
			if !have || maxima[i] > globalMax {
				globalMax = maxima[i]
				have = true
			}
		}
	}
	if !have {
		return 0
	}
	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	fmt.Printf("Сгенерировано элементов: %d\n", len(data))

	fmt.Println("Ищем максимальное значение в один поток")
	start1 := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start1).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start2 := time.Now()
	maxAsync := maxChunks(data)
	elapsedAsync := time.Since(start2).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxAsync, elapsedAsync)
}
