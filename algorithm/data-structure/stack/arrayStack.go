package stack

/*
	栈的常用操作

push 元素入栈
pop  栈顶元素出栈
peek 访问栈元素
*/
type ArrayStack struct {
	data []int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{data: []int{}}
}

func (s *ArrayStack) Push(ele int) {
	s.data = append(s.data, ele)
}

func (s *ArrayStack) Pop() int {
	ret := s.Peek()
	s.data = s.data[:len(s.data)-1]
	return ret
}

func (s *ArrayStack) Peek() int {
	if len(s.data) < 1 {
		return -1
	}
	ret := s.data[len(s.data)-1]
	return ret
}
