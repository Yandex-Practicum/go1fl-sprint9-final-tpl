package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000 // в задаче 10 миллионов
	CHUNKS = 8
)

// считаю что возможность выбора типа была бы уместна uint подойдет оптимальнее и избавит от ряда преобразований на положительный знак

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	data := []int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // time.Now().UnixNano() - чтобы точка отсчета была разной и не выдавало в ответе одно и то же значение

	for i := 0; i < size; i++ {
		data = append(data, r.Intn(size)) // Генерируем случайные числа и добавляем их в срез. Intn - возвращает только положительные значения
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	size := len(data)
	results := make([]int, CHUNKS)

	if size == 0 {
		return 0
	}

	var wg sync.WaitGroup
	rangeSize := size / CHUNKS

	for step := 0; step < CHUNKS; step++ {
		wg.Add(1)

		go func(i int, dataPart []int) {
			defer wg.Done()
			results[i] = maximum(dataPart)
		}(step,
			data[(step*rangeSize):((step+1)*rangeSize)])
	}
	wg.Wait()
	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Printf("\nИщем максимальное значение в один поток\n")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("\nИщем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

}
