package stack

type Elem struct {
	Data string
}

type Stack struct{
	Elems []Elem
	Top int
}

func (s *Stack)Push(e Elem) {
	s.Top++
	s.Elems[s.Top] = e
}

func (s *Stack)Pop() Elem {
	e := s.Elems[s.Top]
	s.Top--
	return e
}

func (s *Stack)IsEmpty() bool {
	return s.Top == 0
}

func (s *Stack)IsFull() bool {
	return s.Top == len(s.Elems) - 1
}

func (s *Stack)GetTop() Elem {
	return s.Elems[s.Top]
}
