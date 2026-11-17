package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	if len(generateRandomElements(0)) != 0 {
		t.Error("должно быть пусто")
	}
	if len(generateRandomElements(-5)) != 0 {
		t.Error("должно быть пусто")
	}
	if len(generateRandomElements(100)) != 100 {
		t.Error("неверная длина")
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{}) != 0 {
		t.Error("пустой слайс")
	}
	if maximum([]int{5}) != 5 {
		t.Error("один элемент")
	}
	if maximum([]int{1, 3, 7, 2, 9, 4}) != 9 {
		t.Error("максимум не найден")
	}
}
