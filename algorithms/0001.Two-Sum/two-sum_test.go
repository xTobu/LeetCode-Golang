package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
	target int
}

type ans struct {
	nums []int
}

type tc struct {
	input input
	a     ans
}

var tcs = []tc{
	tc{
		input: input{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		a: ans{
			nums: []int{1, 2},
		},
	},
	tc{
		input: input{
			nums:   []int{3, 2, 4},
			target: 8,
		},
		a: ans{
			nums: nil,
		},
	},
}

func Test_Problem0001(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		a, input := tc.a, tc.input
		ast.Equal(a.nums, twoSum(input.nums, input.target), "Input:%v", input)
	}
}

func Benchmark_Problem0001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			twoSum(input.nums, input.target)
		}
	}
}
