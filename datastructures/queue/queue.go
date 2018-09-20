package queue

type Elem struct {
	Data string
}

type Queue struct {
	Front int
	Rear int
	Elems []Elem
}

func(q *Queue) IsEmpty() bool {

}

func(q *Queue) IsFull() bool {

}

func(q *Queue) Push(e Elem) {

}

func(q *Queue) Pop() Elem {

}
