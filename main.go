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
		return nil
	}

	src := rand.NewSource(time.Now().Unix())
	data := make([]int, size)
	for i := range size {
		randomNumber := src.Int63()
		data[i] = int(randomNumber%100_000_000 + 1)
		//fmt.Println(data[i])
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	m := data[0]
	for _, v := range data[1:] {
		m = max(m, v)
	}
	return m
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	oneChunk := len(data) / CHUNKS
	if len(data)%CHUNKS != 0 {
		oneChunk++
	}
	sliceMaxValues := make([]int, CHUNKS)
	var muSliceMaxValues sync.Mutex

	var wg sync.WaitGroup
	threads := min(len(data), CHUNKS)
	wg.Add(threads)

	for i := range threads {
		first := i * oneChunk
		last := min(first+oneChunk, len(data))

		go func(chnum int, slice []int) {
			defer wg.Done()

			m := maximum(slice)

			muSliceMaxValues.Lock()
			defer muSliceMaxValues.Unlock()
			sliceMaxValues[chnum] = m
		}(i, data[first:last])
	}

	wg.Wait()
	return maximum(sliceMaxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	// в тз написано микросекунды, но рыба задания была под ms, исправила ms на µs,
	//где опечатка не знаю. посоветовалась с друг говорит µs почти не используется, но оставлю его по тз спринта
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	// С НОВЫМ ГОДОМ, дорогой ментор <3

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
