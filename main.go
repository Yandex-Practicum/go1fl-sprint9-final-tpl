package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Intn(1000) + 1
	}
	return res
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(slice []int, index int) {
			defer wg.Done()
			maxValues[index] = maximum(slice)
		}(data[start:end], i)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	data := generateRandomElements(SIZE)

	start := time.Now()
	maxSingle := maximum(data)
	fmt.Printf("Один поток: %d мкс\n", time.Since(start).Microseconds())

	start = time.Now()
	maxMulti := maxChunks(data)
	fmt.Printf("Восемь потоков: %d мкс\n", time.Since(start).Microseconds())

	if maxSingle == maxMulti {
		fmt.Println("Ок")
	}
}
