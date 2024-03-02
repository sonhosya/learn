package stack

import (
	"reflect"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	stack := NewArrayStack()
	if !reflect.DeepEqual(stack.data, []int{}) {
		t.Fatalf("NewArrayStack err\n")
	}
}

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	stackTarget := &ArrayStack{data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	t.Logf("stack: %v\n", stack)
	t.Logf("stackTarget: %v\n", stackTarget)
	if !reflect.DeepEqual(stack, stackTarget) {
		t.Fatalf("Push err\n")
	}
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	tests := []struct {
		name       string
		arrayStack *ArrayStack
		want       int
	}{
		// TODO: Add test cases.
		{"", stack, 4},
		{"", stack, 3},
		{"", stack, 2},
		{"", stack, 1},
		{"", stack, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arrayStack.Pop(); got != tt.want {
				t.Errorf("ArrayStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Peek(t *testing.T) {
	stack := NewArrayStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	if stack.Peek() != 9 {
		t.Fatalf("Push err\n")
	}
}
