package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{
			name:    "zero size",
			size:    0,
			wantLen: 0,
		},
		{
			name:    "negative size",
			size:    -10,
			wantLen: 0,
		},
		{
			name:    "positive size",
			size:    100,
			wantLen: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d) length = %d, want %d", tt.size, len(got), tt.wantLen)
			}

			for _, value := range got {
				if value <= 0 {
					t.Errorf("generateRandomElements(%d) returned non-positive value: %d", tt.size, value)
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
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "maximum in middle",
			data: []int{1, 2, 100, 3, 4},
			want: 100,
		},
		{
			name: "maximum first",
			data: []int{100, 1, 2, 3},
			want: 100,
		},
		{
			name: "maximum last",
			data: []int{1, 2, 3, 100},
			want: 100,
		},
		{
			name: "duplicates",
			data: []int{7, 7, 7, 7},
			want: 7,
		},
		{
			name: "negative numbers",
			data: []int{-10, -3, -20},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
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
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "less than chunks",
			data: []int{3, 1, 2},
			want: 3,
		},
		{
			name: "exactly chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "more than chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 10,
		},
		{
			name: "maximum in last chunk",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
			want: 100,
		},
		{
			name: "duplicates",
			data: []int{5, 5, 5, 5, 5, 5, 5, 5, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}