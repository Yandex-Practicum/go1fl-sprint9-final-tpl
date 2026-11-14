package main

import (
	"strconv"
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, 0, 1, 10, 1000} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			elements := generateRandomElements(size)
			expectedLength := size
			if expectedLength < 0 {
				expectedLength = 0
			}
			if len(elements) != expectedLength {
				t.Fatalf("generateRandomElements(%d): длина %d, ожидалось %d", size, len(elements), expectedLength)
			}
			for i, value := range elements {
				if value <= 0 {
					t.Fatalf("элемент %d: ожидалось положительное число, получено %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "one", data: []int{7}, want: 7},
		{name: "ascending", data: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "descending", data: []int{9, 6, 4, 1}, want: 9},
		{name: "middle", data: []int{3, 12, 1, 7}, want: 12},
		{name: "equal", data: []int{5, 5, 5}, want: 5},
		{name: "zero", data: []int{0, 0, 0}, want: 0},
		{name: "negative", data: []int{-11, -3, -9}, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "one", data: []int{7}, want: 7},
		{name: "less than eight", data: []int{9, 3, 5, 2, 1}, want: 9},
		{name: "eight", data: []int{1, 2, 3, 40, 5, 6, 7, 8}, want: 40},
		{name: "remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 10}, want: 99},
		{name: "non-last chunk", data: []int{1, 4, 70, 3, 9, 6, 8, 2, 5, 10, 11, 12, 13, 14, 15, 16}, want: 70},
		{name: "all negative", data: []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
		{name: "zeroes", data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}

	data := make([]int, 10_001)
	for i := range data {
		data[i] = i
	}
	data[5_000] = 100_000
	if got := maxChunks(data); got != maximum(data) {
		t.Errorf("maxChunks() = %d, maximum() = %d", got, maximum(data))
	}
}
