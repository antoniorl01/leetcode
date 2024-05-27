package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stk := Stack{}

	for _, i := range s {
		if stk.IsEmpty() == true {
			stk.Push(byte(i))
		} else {
			acc := pusheable(stk.Peek(), byte(i))

			switch acc {
			case 0:
				return false
			case 1:
				stk.Push(byte(i))
			case 2:
				stk.Pop()
			default:
				return false
			}
		}
	}

	return stk.IsEmpty()
}

// 0 -> Err
// 1 -> Push
// 2 -> Pop
func pusheable(peak, waiting byte) int {
	if peak == '[' && (waiting == '[' || waiting == '(' || waiting == '{' || waiting == ']') {
		if waiting == ']' {
			return 2
		}

		return 1
	}

	if peak == '{' && (waiting == '[' || waiting == '(' || waiting == '{' || waiting == '}') {
		if waiting == '}' {
			return 2
		}

		return 1
	}

	if peak == '(' && (waiting == '[' || waiting == '(' || waiting == '{' || waiting == ')') {
		if waiting == ')' {
			return 2
		}

		return 1
	}

	return 0
}

type Stack struct {
	items []byte
}

func (s *Stack) Push(item byte) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() byte {
	if len(s.items) == 0 {
		panic("El stack está vacío")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

func (s *Stack) Peek() byte {
	if len(s.items) == 0 {
		panic("El stack está vacío")
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
