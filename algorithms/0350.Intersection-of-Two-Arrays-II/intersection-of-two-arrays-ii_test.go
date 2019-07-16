package problem0350

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	nums1 []int
	nums2 []int
}

type output struct {
	answer []int
}

var tcs = []tc{

	tc{
		input{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
		},
		output{answer: []int{2, 2}},
	},

	tc{
		input{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
		},
		output{answer: []int{4, 9}},
	},
}

func Test_Problem0350(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, intersect(input.nums1, input.nums2), "Input: %v", input)

	}
}

func Test_Problem0350_2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, intersect2(input.nums1, input.nums2), "Input: %v", input)

	}
}

func Benchmark_Problem0350(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			intersect(input.nums1, input.nums2)
		}
	}
}

func Benchmark_Problem0350_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			intersect2(input.nums1, input.nums2)
		}
	}
}
