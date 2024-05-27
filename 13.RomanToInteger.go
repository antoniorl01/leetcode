package main

func romanToInt(s string) int {
	var num int

	for idx := len(s) - 1; idx >= 0; idx-- {
		curr := cast(s[idx])

		if idx != len(s)-1 {
			prev := cast(s[idx+1])

			if curr < prev {
				num -= curr
			} else {
				num += curr
			}
		} else {
			num += curr
		}
	}

	return num
}

func cast(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
