package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	// размер 0 -> пустой слайс
	if got := generateRandomElements(0); len(got) != 0 {
		t.Errorf("expected empty slice for size 0, got %v", got)
	}
	// размер -1 -> пустой слайс
	if got := generateRandomElements(-10); len(got) != 0 {
		t.Errorf("expected empty slice for negative size, got %v", got)
	}
	// размер > 0
	size := 1000
	data := generateRandomElements(size)
	if len(data) != size {
		t.Errorf("expected length %d, got %d", size, len(data))
	}
	// Все элементы > 0
	for _, v := range data {
		if v <= 0 {
			t.Errorf("expected positive number, got %d", v)
		}
	}
}

func TestMaximum(t *testing.T) {
	// пустой слайс
	if got := maximum([]int{}); got != 0 {
		t.Errorf("expected 0 for empty slice, got %d", got)
	}
	// слайс с одним элементом
	if got := maximum([]int{42}); got != 42 {
		t.Errorf("expected 42, got %d", got)
	}
	// слайс с несколькими элементами
	data := []int{1, 3, 2, 5, 4}
	if got := maximum(data); got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
	// слайс с одинаковыми элементами
	data = []int{7, 7, 7}
	if got := maximum(data); got != 7 {
		t.Errorf("expected 7, got %d", got)
	}
}
