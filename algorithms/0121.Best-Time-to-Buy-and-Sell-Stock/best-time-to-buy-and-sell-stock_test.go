package problem0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	prices []int
}

type output struct {
	profit int
}

var tcs = []tc{
	tc{
		input{
			prices: []int{7, 1, 5, 3, 6, 4},
		},
		output{
			profit: 5,
		},
	},
	tc{
		input{
			prices: []int{7, 6, 4, 3, 1},
		},
		output{
			profit: 0,
		},
	},
}

func Test_Problem0121(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		input := tc.input
		output := tc.output
		ast.Equal(output.profit, maxProfit(input.prices), "Input: %v", input)

	}
}

func Benchmark_Problem0121(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			maxProfit(input.prices)
		}
	}
}
