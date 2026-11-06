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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for index := range data {
		data[index] = rand.Intn(1_000_000) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkCount := min(CHUNKS, len(data))
	chunkSize := len(data) / chunkCount
	maxima := make([]int, chunkCount)
	var workers sync.WaitGroup

	workers.Add(chunkCount)
	for index := range chunkCount {
		start := index * chunkSize
		end := start + chunkSize
		if index == chunkCount-1 {
			end = len(data) // В последнюю часть войдёт и остаток.
		}
		go func(index int, chunk []int) {
			defer workers.Done()
			maxima[index] = maximum(chunk)
		}(index, data[start:end])
	}

	workers.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Последовательный поиск")
	start := time.Now()
	largest := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)

	fmt.Printf("Поиск в %d горутинах\n", CHUNKS)
	start = time.Now()
	largest = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)
}
