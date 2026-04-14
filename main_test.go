package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "Positive count",
			input: 1000,
			want:  1000,
		},
		{name: "Zero elements",
			input: 0,
			want:  0,
		},
		{name: "Negative count",
			input: -10,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.input)
			assert.Len(t, got, tt.want)
			//дополнительная проверка, что слайс состоит только из положительных чисел
			if len(got) > 0 {
				for _, v := range got {
					require.GreaterOrEqual(t, v, 0)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "Empty slice",
			input: make([]int, 0),
			want:  0,
		},
		{name: "Nil slice",
			input: nil,
			want:  0,
		},
		{name: "One element in slice",
			input: []int{5},
			want:  5,
		},
		//Несмотря на то что функция GenerateRandomElements отдает слайс исключительно из положительных
		//чисел, все равно нужна проверка на поиск максимального числа из отрицательных чисел
		{name: "Negative elements in slice",
			input: []int{-10, -9, -8, -20, -1},
			want:  -1,
		},
		{name: "All elements are equal",
			input: []int{7, 7, 7, 7},
			want:  7,
		},
		{name: "The greater number is in the end",
			input: []int{7, 6, 7, 10},
			want:  10,
		},
		{name: "The greater number is in the beginning",
			input: []int{17, 6, 7, 10},
			want:  17,
		},
		{name: "The greater number is in the middle of the slice",
			input: []int{17, 6, 23, 10, 22},
			want:  23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
