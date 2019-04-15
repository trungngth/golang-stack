package stack

//Stack implement
type Stack struct {
	top    *node
	length int
}
type node struct {
	value string
	prev  *node
}

// New return a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len return length of the stack
func (s *Stack) Len() int {
	return s.length
}

// Peek return the top item on the stack
func (s *Stack) Peek() string {
	if s.length == 0 {
		return ""
	}
	return s.top.value
}

// Pop return the top item of the stack and return it
func (s *Stack) Pop() string {
	if s.length == 0 {
		return ""
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push return a value onto the top of the stack
func (s *Stack) Push(value string) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
