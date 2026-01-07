package main

// Пишите тесты в этом файле

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(0)
	if len(data) != 0 {
		t.Error("expected 0 lenght slice")
	}

	data = generateRandomElements(3)
	if len(data) != 3 {
		t.Error("wrong slice lenght")
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{}) != 0 {
		t.Error("expected 0 for an empty slice")
	}

	if maximum([]int{5}) != 5 {
		t.Error("wrong one element maximum")
	}

	if maximum([]int{1, 5, 2, 3, 4}) != 5 {
		t.Error("wrong maximum number")
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{1, 5, 2, 3, 4}
	m := maxChunks(data)
	if m != 5 {
		t.Error("wrong maximum number in chunks", "expected", 5, "receieved", m)
	}

	if maxChunks([]int{}) != 0 {
		t.Error("expected 0 for an empty slice")
	}
}
