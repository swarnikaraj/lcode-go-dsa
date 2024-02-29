package main

import "strings"

func reverseWords(s string) string {
	trimmed := strings.TrimSpace(s)
	arr := strings.Split(trimmed, " ")

	newstr := ""

	for i := len(arr) - 1; i > 0; i-- {
		w := strings.TrimSpace(arr[i])
		if len(w) > 0 {
			newstr = newstr + arr[i] + " "
		}

	}
	newstr = newstr + strings.TrimSpace(arr[0])
	return strings.TrimSpace(newstr)
}
func main() {
	s := "a good   example"
	reverseWords(s)
}