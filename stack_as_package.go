package stack_as_package

type Stack struct {
	i int
	data [10]int
}

func (stack *Stack) Push(x int) {
	stack.data[stack.i] = x
	stack.i++
}

func (stack *Stack) Pop() (ret int) {
	stack.i--
	ret = stack.data[stack.i]
	return
}