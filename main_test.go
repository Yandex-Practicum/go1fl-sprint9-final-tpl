package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nilSlice", []int{}, 0},
		{"oneElem", []int{1}, 1},
		{"equalElem", []int{1, 1, 1, 1}, 1},
		{"maxStart", []int{100, 2, 3, 4}, 100},
		{"maxEnd", []int{1, 2, 3, 999}, 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
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
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 3, 2, 5, 4}, 5},
		{"All equal", []int{7, 7, 7}, 7},
		{"Mixed numbers", []int{1, 0, 10, 5}, 10},
		{"Nil slice", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
