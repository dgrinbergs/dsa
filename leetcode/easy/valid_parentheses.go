package easy

// https://neetcode.io/problems/validate-parentheses

var opposites = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

func validParentheses(s string) bool {
	// odd length strings cannot have equal numbers of opening and closing parens
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)

	for _, ch := range s {
		// add all the opening parens
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		// check if the closing bracket for most recent opening one matches the current rune
		if opposites[stack[len(stack)-1]] != ch {
			return false
		}

		// 'pop' it off the stack
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
