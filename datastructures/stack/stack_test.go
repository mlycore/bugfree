package stack
import (
	"fmt"
	"testing"
)

func Test_Balance(t *testing.T) {
	// can't handle more right brackets
	b := []string{"{", "{", "[", "]", "}", "}"}
	s := &Stack{
		Elems: make([]Elem, len(b)),
		Top: -1,
	}

	for _, i := range b {
		if i == "{" || i == "[" {
			s.Push(Elem{Data: i})
			fmt.Printf("%v, %d\n", s.Elems, s.Top)
		} 
		if i == "}" {
			if s.GetTop().Data == "{" {
				s.Pop()
			}
		}
		if i == "]" {
			if s.GetTop().Data == "[" {
				s.Pop()
			}
		}
	}
	fmt.Printf("%v, %d\n", s.Elems, s.Top)
}
