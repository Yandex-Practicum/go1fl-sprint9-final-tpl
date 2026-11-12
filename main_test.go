package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-5, 0, 1, 1000} {
		data := generateRandomElements(size)

		wantLen := max(size, 0)
		if len(data) != wantLen {
			t.Fatalf("size %d: got len %d, want %d", size, len(data), wantLen)
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("size %d: got non-positive value %d", size, value)
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
		{"one element", []int{7}, 7},
		{"max first", []int{9, 1, 2}, 9},
		{"max last", []int{1, 2, 9}, 9},
		{"duplicates", []int{5, 5, 5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
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
		{"empty", []int{}, 0},
		{"fewer than CHUNKS", []int{3, 1, 2}, 3},
		{"max in remainder tail", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 42}, 42},
		{"max first", []int{42, 1, 1, 1, 1, 1, 1, 1, 1}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}

	data := generateRandomElements(10_000)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Errorf("maxChunks = %d, maximum = %d", got, want)
	}
}
