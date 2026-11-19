package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(100)
	if len(data) != 100 {
		t.Errorf("Ожидали 100 чисел, получили %d", len(data))
	}
	for _, number := range data {
		if number <= 0 {
			t.Errorf("Число должно быть положительным, получили %d", number)
		}
	}

	if len(generateRandomElements(0)) != 0 {
		t.Error("При размере 0 слайс должен быть пустым")
	}
	if len(generateRandomElements(-5)) != 0 {
		t.Error("При отрицательном размере слайс должен быть пустым")
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{3, 10, 5, 2}) != 10 {
		t.Error("Неверный максимум для нескольких чисел")
	}
	if maximum([]int{7}) != 7 {
		t.Error("Неверный максимум для одного числа")
	}
	if maximum([]int{}) != 0 {
		t.Error("Для пустого слайса ожидается 0")
	}
	if maximum([]int{5, 5, 5}) != 5 {
		t.Error("Неверный максимум для одинаковых чисел")
	}
	if maximum([]int{-4, -2, -8}) != -2 {
		t.Error("Неверный максимум для отрицательных чисел")
	}
}

func TestMaxChunks(t *testing.T) {
	if maxChunks([]int{}) != 0 {
		t.Error("Для пустого слайса ожидается 0")
	}
	if maxChunks([]int{15}) != 15 {
		t.Error("Неверный максимум для одного числа")
	}
	if maxChunks([]int{1, 3, 7, 2, 5}) != 7 {
		t.Error("Неверный максимум для слайса короче восьми элементов")
	}
	if maxChunks([]int{4, 2, 9, 3, 5, 7, 1, 6, 8, 20}) != 20 {
		t.Error("Неверный максимум, когда длина не делится на 8")
	}
}
