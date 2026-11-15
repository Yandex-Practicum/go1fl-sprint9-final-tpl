package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "zero size", size: 0, want: 0},
		{name: "negative size", size: -10, want: 0},
		{name: "one element", size: 1, want: 1},
		{name: "small slice", size: 100, want: 100},
		{name: "big slice", size: 10_000, want: 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.want {
				t.Fatalf("generateRandomElements(%d): длина = %d, ожидалось %d",
					tt.size, len(got), tt.want)
			}
			if got == nil {
				t.Fatalf("generateRandomElements(%d): вернулся nil вместо слайса", tt.size)
			}
			for i, v := range got {
				if v <= 0 {
					t.Fatalf("generateRandomElements(%d): элемент %d не положительный: %d",
						tt.size, i, v)
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
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "all equal", data: []int{7, 7, 7, 7}, want: 7},
		{name: "max in the middle", data: []int{1, 5, 9, 3, 2}, want: 9},
		{name: "max at the start", data: []int{100, 50, 20, 1}, want: 100},
		{name: "max at the end", data: []int{3, 4, 5, 6, 99}, want: 99},
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
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{15}, want: 15},
		{name: "less than CHUNKS", data: []int{4, 8, 2}, want: 8},
		{name: "exactly CHUNKS", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "more than CHUNKS", data: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, want: 100},
		{name: "max in last chunk", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 999}, want: 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaximumAndMaxChunksAgree(t *testing.T) {
	data := generateRandomElements(100_000)
	if maximum(data) != maxChunks(data) {
		t.Errorf("maximum и maxChunks вернули разные значения")
	}
}
