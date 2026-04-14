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
		return make([]int, 0)
	}
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return data[0]
	}

	maxNum := data[0]
	for _, v := range data[1:] {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return data[0]
	}

	wg := sync.WaitGroup{}

	findingMax := func(data []int) int {
		maxNum := data[0]
		for _, v := range data[1:] {
			if v > maxNum {
				maxNum = v
			}
		}
		return maxNum
	}

	maxSlice := make([]int, CHUNKS)
	sizeOfSlice := len(data) / CHUNKS

	//Кейс если длина слайса меньше кол-ва CHUNKS
	if sizeOfSlice == 0 {
		return findingMax(data)
	}

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		startIndex := i * sizeOfSlice
		var endIndex int

		// На случай, если число элементов слайса не будет делиться нацело
		if i != CHUNKS-1 {
			endIndex = startIndex + sizeOfSlice
		} else {
			endIndex = len(data)
		}
		chSlice := data[startIndex:endIndex]

		go func(chunkSlice []int, index int) {
			defer wg.Done()
			maxSlice[index] = findingMax(chunkSlice)
		}(chSlice, i)
	}
	wg.Wait()

	return findingMax(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slice)
	// по условию задания надо выводить в микросекундах
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	// по условию задания надо выводить в микросекундах
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
