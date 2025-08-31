package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testGRE struct {
	size int
	exp  int
}

type testMax struct {
	name string
	data []int
	exp  int
}

// проверяем, что функция вернет нужное количество элементов
func TestGenerateRandomElements(t *testing.T) {
	vals := []testGRE{{0, 0}, {-1, 0}, {999, 999}}

	for _, v := range vals {
		t.Run("test "+strconv.Itoa(v.size), func(t *testing.T) {
			slice := generateRandomElements(v.size)

			require.NotNil(t, slice)
			assert.Equal(t, v.exp, len(slice))

		})
	}
}

// проверяем, что функция работает с такими случаями как пустой слайс, слайс с одним элементом и так далее.
func TestMaximum(t *testing.T) {
	datas := []testMax{
		{"Nill test", nil, 0},
		{"Empty test", []int{}, 0},
		{"Alone test", []int{99}, 99},
		{"Twins test", []int{999, 999, 999}, 999},
		{"Regular test", []int{10, 99999, 30, 10, 55, 123, 654, 8484484}, 8484484},
	}

	for _, d := range datas {
		t.Run(d.name, func(t *testing.T) {
			result := maximum(d.data)
			assert.Equal(t, d.exp, result)
		})
	}
}
