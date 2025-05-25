package main

import (
	"testing"
)

// Тесты для функции maximum()
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Все элементы равны", []int{7, 7, 7, 7}, 7},
		{"Максимум в начале", []int{100, 2, 3, 4}, 100},
		{"Максимум в конце", []int{1, 2, 3, 999}, 999},
		{"Средний максимум", []int{1, 5000, 3, 4, 2}, 5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

// Тесты для функции generateRandomElements()
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Нулевой размер", 0},
		{"Отрицательный размер", -5},
		{"Маленький размер", 5},
		{"Средний размер", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if tt.size <= 0 && len(result) != 0 {
				t.Errorf("generateRandomElements(%d) должен вернуть пустой слайс, получил длину %d", tt.size, len(result))
			}
			if tt.size > 0 {
				if len(result) != tt.size {
					t.Errorf("generateRandomElements(%d) вернул слайс длины %d, ожидается %d", tt.size, len(result), tt.size)
				}
				for i, v := range result {
					if v < 1 {
						t.Errorf("элемент %d равен %d, должен быть >= 1", i, v)
					}
				}
			}
		})
	}
}
