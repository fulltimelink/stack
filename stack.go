package stack

import "errors"

// Stack --  @# 外部封装类型
type Stack[T any] struct {
	stackSlice []T
}

// --  @# 内部封装类型
type stackSlice[T any] []T

// ErrorStackOverflow --  @# 哨兵异常定义
var ErrorStackOverflow = errors.New("overflow")

// Pop --  @# 出栈
func (s *Stack[T]) Pop() (ret T, err error) {
	l := len(s.stackSlice)
	if 0 < l {
		ret = (s.stackSlice)[l-1]
		s.stackSlice = (s.stackSlice)[:l-1]
		return
	}
	err = ErrorStackOverflow
	return
}

// Push --  @# 入栈
func (s *Stack[T]) Push(e T) {
	s.stackSlice = append(s.stackSlice, e)
}

// Len --  @# 栈长
func (s *Stack[T]) Len() int {
	return len(s.stackSlice)
}

// IsEmpty --  @# 栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return 0 == s.Len()
}

// NewStack --  @# 初始化一个栈
func NewStack[T any]() Stack[T] {
	return Stack[T]{stackSlice[T]{}}
}
