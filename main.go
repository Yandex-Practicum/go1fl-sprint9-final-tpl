package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000 // Размер массива (10 миллионов элементов)
	CHUNKS = 8          // Количество частей и горутин
)

// generateRandomElements создает слайс заданного размера и заполняет его случайными положительными числами.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1_000_000) + 1 // Генерируем числа от 1 до 1 000 000
	}
	return slice
}

// maximum находит максимальное число в слайсе.
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

// maxChunks находит максимальное число, разделяя слайс на CHUNKS частей и используя горутины.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	// Если элементов меньше, чем частей, просто ищем максимум обычным способом
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Последняя горутина забирает "хвост", если длина не делится на 8 нацело
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Запускаем горутину для поиска максимума в куске
		go func(idx, s, e int) {
			defer wg.Done()
			maxes[idx] = maximum(data[s:e])
		}(i, start, end)
	}

	// Ждем завершения всех 8 горутин
	wg.Wait()

	// Ищем максимум среди 8 найденных максимумов
	return maximum(maxes)
}

func main() {
	fmt.Println("Генерация данных...")
	data := generateRandomElements(SIZE)

	// 1. Поиск в один поток
	start := time.Now()
	maxSingle := maximum(data)
	durationSingle := time.Since(start)

	// 2. Поиск в 8 потоков
	start = time.Now()
	maxMulti := maxChunks(data)
	durationMulti := time.Since(start)

	// Вывод результатов
	fmt.Printf("Размер массива: %d элементов\n", SIZE)
	fmt.Printf("Максимум (1 поток): %d, Время: %d мкс\n", maxSingle, durationSingle.Microseconds())
	fmt.Printf("Максимум (8 потоков): %d, Время: %d мкс\n", maxMulti, durationMulti.Microseconds())
}