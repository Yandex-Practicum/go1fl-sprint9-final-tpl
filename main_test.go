package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"zero size", 0, 0},
		{"negative size", -7, 0},
		{"size 1", 1, 1},
		{"size 100", 100, 100},
		{"size 1000", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Fatalf("len = %d, want %d", len(got), tt.wantLen)
			}
			for i, v := range got {
				if v < 0 {
					t.Errorf("element %d is negative: %d", i, v)
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
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"single zero", []int{0}, 0},
		{"all equal", []int{5, 5, 5, 5}, 5},
		{"max at start", []int{9, 1, 2, 3}, 9},
		{"max at end", []int{1, 2, 3, 9}, 9},
		{"max in middle", []int{1, 2, 9, 3}, 9},
		{"two elements", []int{3, 7}, 7},
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
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"shorter than CHUNKS", []int{3, 1, 4}, 4},
		{"exactly CHUNKS elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"many elements", []int{5, 1, 9, 2, 8, 3, 7, 4, 6, 0, 11, 10}, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
