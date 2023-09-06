package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	openedBrackets := []rune{}
	matchedBrackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	countOpen := 0
	countClose := 0

	for _, char := range s {
		if char == '{' ||
			char == '[' ||
			char == '(' {
			openedBrackets = append(openedBrackets, char)
			countOpen++
		} else {
			countClose++
			if len(openedBrackets) > 0 && openedBrackets[len(openedBrackets)-1] == matchedBrackets[char] {
				// delete matched opened brackets
				openedBrackets = openedBrackets[:len(openedBrackets)-1]
			}
		}
	}
	return len(openedBrackets) == 0 && countClose == countOpen

}

func main() {
	target := isValid("(") // false
	fmt.Printf("\n(: %v", target)
	target = isValid("]") // false
	fmt.Printf("\n]: %v", target)
	target = isValid("([}}])") // false
	fmt.Printf("\n([}}]): %v", target)
	target = isValid("()") // true
	fmt.Printf("\n(): %v", target)
	target = isValid("()[]{}") // true
	fmt.Printf("\n()[]{}: %v", target)
	target = isValid("(]") // false
	fmt.Printf("\n(]: %v", target)
}
