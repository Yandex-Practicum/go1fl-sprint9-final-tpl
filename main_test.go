package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 32} {
		data := generateRandomElements(size)
		if len(data) != max(size, 0) {
			t.Fatalf("generateRandomElements(%d): длина %d", size, len(data))
		}
		for _, value := range data {
			if value <= 0 {
				t.Errorf("generateRandomElements(%d): получено %d, нужно положительное число", size, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"negative", []int{-8, -3, -11}, -3},
		{"repeated", []int{4, 9, 1, 9}, 9},
		{"eight chunks", []int{-1, -2, -3, -4, -5, -6, -7, -8}, -1},
		{"middle chunk", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 42, 11, 12, 13, 14, 15, 16}, 42},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 99}, 99},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum() = %d, want %d", got, test.want)
			}
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks() = %d, want %d", got, test.want)
			}
		})
	}
}
