package main

type Stack struct {
	ops []string
}

func (s *Stack) Push(v string) {
	s.ops = append(s.ops, v)
}

func (s *Stack) Pop() string {
	var remove string
	if len(s.ops) > 0 {
		remove = s.ops[len(s.ops)-1]
		s.ops = s.ops[:len(s.ops)-1]
	}
	return remove
}

type Queue struct {
	RPN []string
}

func (q *Queue) Enqueue(v string) {
	q.RPN = append(q.RPN, v)
}
