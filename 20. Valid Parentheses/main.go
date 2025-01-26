package main

import "fmt"

type stack struct {
	size int
	data []string
}

func (s *stack) Create() {
	s.size = 0
	s.data = []string{}
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) Push(value string) {
	s.data = append(s.data, value)
	s.size++
}

func (s *stack) Pop() (string, bool) {
	if s.size == 0 {
		return "", false
	}
	value := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return value, true
}

func (s *stack) Top() string {
	if s.size != 0 {
		return s.data[s.size-1]
	} else {
		return ""
	}
}

func isValid(s string) bool {
	var bigStack stack
	bigStack.Create()
	for _, i := range s {
		if string(i) == "(" {
			bigStack.Push(")")
		}
		if string(i) == "{" {
			bigStack.Push("}")
		}
		if string(i) == "[" {
			bigStack.Push("]")
		}
		if string(i) == ")" || string(i) == "]" || string(i) == "}" {
			if bigStack.Top() != string(i) {
				return false
			}
			bigStack.Pop()
		}
	}
	if bigStack.Size() != 0 {
		return false
	}
	return true
}

func main() {
	str := "()"
	if isValid(str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
