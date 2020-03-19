package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true

*/
//
//func main() {
//	fmt.Println(isValid("()"))
//	fmt.Println(isValid("()[]{}"))
//	fmt.Println(isValid("(]"))
//	fmt.Println(isValid("([)]"))
//	fmt.Println(isValid("{[]}"))
//	fmt.Println(isValid("{[(]}"))
//}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	m := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	stack := make([]string, 0)

	for _, e := range s {
		if len(stack) > 0 {
			// 栈顶元素匹配
			if stack[len(stack)-1] == m[string(e)] {
				stack = stack[:len(stack)-1]
				continue
			}
			// 未匹配，且是右括号的，则直接返回
			if _, ok := m[string(e)]; ok {
				return false
			}
		}
		stack = append(stack, string(e))
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
