package stack

import "fmt"

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (s *Stack) Push(n int) {
	if s.ix >= LIMIT {
		return
	}
	s.data[s.ix] = n
	s.ix++
}
func (s *Stack) Pop() (n int) {
	if s.ix != 0 {
		n = s.data[s.ix-1]
		s.data[s.ix-1] = 0
		s.ix--
	}
	return
}
func (s Stack) String() string {
	str := ""
	for i := 0; i < s.ix; i++ {
		str += fmt.Sprintf("[%d:%d] ", i, s.data[i])
	}
	return str
}
