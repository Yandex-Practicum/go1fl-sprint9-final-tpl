package main

// Пишите тесты в этом файле

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"Нормальный размер", 10, 10},
		{"Нулевой размер", 0, 0},
		{"Отрицательный размер", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d) вернул слайс длиной %d, ожидалось %d", tt.size, len(got), tt.want)
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
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Несколько элементов", []int{1, 5, 3, 9, 2}, 9},
		{"Все элементы одинаковые", []int{7, 7, 7, 7}, 7},
		{"Максимум в конце", []int{1, 2, 3, 100}, 100},
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
		{"Пустой слайс", []int{}, 0},
		{"Меньше 8 элементов", []int{1, 5, 2}, 5},
		{"Ровно 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"Больше 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"Максимум в середине", []int{10, 20, 30, 100, 40, 50, 60, 70, 80, 90}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
