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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	slice := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(1_000_000) + 1
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	//обработка ошибок
	if len(data) == 0 {
		return 0
	}

	//создание счетчика
	var wg sync.WaitGroup
	//создания среза для 8 максимальных чисел
	maxElements := make([]int, CHUNKS)
	//создание среза путем деления основного массива
	chunkSize := len(data) / CHUNKS
	//инициализация счетчика
	wg.Add(CHUNKS)
	//цикл с 8-ю итерациями
	for i := range CHUNKS {
		// находим начальный индекс среза
		start := i * chunkSize
		// находим конечный индекс среза
		end := start + chunkSize

		go func(idx, s, e int) {
			defer wg.Done()
			maxElements[idx] = maximum(data[s:e])
		}(i, start, end)
	}
	wg.Wait()
	return maximum(maxElements)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(elements)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
