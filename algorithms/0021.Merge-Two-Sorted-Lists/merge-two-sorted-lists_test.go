package problem0021

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
	l1 []int
	l2 []int
}

type output struct {
	answer []int
}

var tcs = []tc{

	tc{
		input{
			l1: []int{1, 2, 4},
			l2: []int{1, 3, 4},
		},
		output{
			answer: []int{1, 1, 2, 3, 4, 4},
		},
	},
	tc{
		input{
			l1: []int{},
			l2: []int{1, 3, 4},
		},
		output{[]int{1, 3, 4}},
	},
	tc{
		input{
			l1: []int{1, 3, 4},
			l2: []int{},
		},
		output{[]int{1, 3, 4}},
	},
}

func Test_Problem0206(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		l1 := utils.Ints2ListNode(input.l1)
		l2 := utils.Ints2ListNode(input.l2)
		ast.Equal(output.answer, utils.ListNode2Ints(mergeTwoLists(l1, l2)), "Input: %v", input)

	}
}

func Benchmark_Problem0206(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			l1 := utils.Ints2ListNode(input.l1)
			l2 := utils.Ints2ListNode(input.l2)
			utils.ListNode2Ints(mergeTwoLists(l1, l2))
		}
	}
}
