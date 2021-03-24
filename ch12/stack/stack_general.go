package stack

import "errors"

type Stack []interface{}

func (st Stack) Len() int {
	return len(st)
}
func (st Stack) Cap() int {
	return cap(st)
}
func (st Stack) IsEmpty() bool {
	return st.Len() == 0
}
func (st *Stack) Push(e interface{}) {
	*st = append(*st, e)
}
func (st Stack) Top() (interface{}, error) {
	if len(st) == 0 {
		return nil, errors.New("stack is empty")
	}
	return st[len(st)-1], nil
}
func (st *Stack) Pop() (interface{}, error) {
	stk := *st // dereference to a local variable stk
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*st = stk[:len(stk)-1] // shrink the stack
	return top, nil
}
