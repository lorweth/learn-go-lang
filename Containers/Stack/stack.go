package Container

type Node struct {
	Value interface{}
	Next  *Node
}

type Stack struct {
	Top *Node
}

// Add a new element to the top of the stack
func (s *Stack) Push(v interface{}) {
	n := &Node{v, nil}
	n.Next = s.Top
	s.Top = n
}

// Remove the top element from the stack
func (s *Stack) Pop() interface{} {
	if s.Top == nil {
		return nil
	}
	v := s.Top.Value
	s.Top = s.Top.Next
	return v
}

// Return the top element of the stack
func (s *Stack) Peek() interface{} {
	if s.Top == nil {
		return nil
	}
	return s.Top.Value
}

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}
