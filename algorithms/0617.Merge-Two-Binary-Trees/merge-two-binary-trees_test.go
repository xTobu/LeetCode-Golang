package problem0617

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xTobu/LeetCode-Golang/lib/utils"
)

type tc struct {
	input
	output
}

type input struct {
	t1 []int
	t2 []int
}

type output struct {
	answer []int
}

var NULL = -1 << 63
var tcs = []tc{
	tc{
		input{
			t1: []int{1, 3, 2, 5},
			t2: []int{2, 1, 3, NULL, 4, NULL, 7},
		},
		output{
			answer: []int{3, 4, 5, 5, 4, NULL, 7},
		},
	},
	tc{
		input{
			t1: []int{},
			t2: []int{2, 1, 3, NULL, 4, NULL, 7},
		},
		output{
			answer: []int{2, 1, 3, NULL, 4, NULL, 7},
		},
	},
	tc{
		input{
			t1: []int{1, 3, 2, 5},
			t2: []int{},
		},
		output{
			answer: []int{1, 3, 2, 5},
		},
	},
	tc{
		input{
			t1: []int{},
			t2: []int{},
		},
		output{
			answer: []int{},
		},
	},
}

func Test_Problem0617(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		input := tc.input
		output := tc.output
		t1 := utils.IntsToTreeNode(input.t1)
		t2 := utils.IntsToTreeNode(input.t2)
		answer := mergeTrees(t1, t2)
		ast.Equal(utils.IntsToTreeNode(output.answer), answer, "Input :%v", input)

	}
}

func Benchmark_Problem0617(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			t1 := utils.IntsToTreeNode(input.t1)
			t2 := utils.IntsToTreeNode(input.t2)
			mergeTrees(t1, t2)
		}
	}
}
