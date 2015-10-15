package stack_as_package_test

import (
	"testing"
	"github.com/novieq/stack-as-package"
)
func TestPushPop(t *testing.T) {
	c := new(stack_as_package.Stack)
	c.Push(5)
	ret := c.Pop()
	if ret != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	}
}
