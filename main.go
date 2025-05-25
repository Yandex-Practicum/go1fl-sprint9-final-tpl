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

// generateRandomElements генерирует слайс с random положительными числами.
// Если size <= 0, возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	rand.Seed(time.Now().UnixNano())
	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(1_000_000) + 1 // положительные числа от 1 до 1_000_000
	}
	return elements
}

// maximum возвращает максимум из слайса или 0, если слайс пуст.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks делит слайс на CHUNKS частей и ищет максимум в каждой части параллельно.
func maxChunks(data []int) int {
	length := len(data)
	if length == 0 {
		return 0
	}

	chunkSize := length / CHUNKS
	maximums := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = length
		}

		go func(idx, s, e int) {
			defer wg.Done()
			maximums[idx] = maximum(data[s:e])
		}(i, start, end)
	}

	wg.Wait()
	return maximum(maximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	arr := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(arr)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(arr)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
