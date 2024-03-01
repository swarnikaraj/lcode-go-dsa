package main

import (
	"fmt"
	"strings"
)

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

func rotateString(s string, goal string) bool {
 if len(s) != len(goal) {
        return false
    }
    
    if s == goal {
        return true
    }
    finstr:=s
    for i:=0;i<len(s);i++{
        leftmost:=string(finstr[0])
        fmt.Println(leftmost)

       substr:=finstr[1:]
       fmt.Println(substr)
       joinstr:=substr + leftmost
       fmt.Println(joinstr)
       if joinstr == goal{
           return true
       }
        finstr=joinstr
        
    }

    return finstr==goal
}
func main() {
	s := "a good   example"
	reverseWords(s)
s1:= "abcde"
 goal := "cdeab"
	rotateString(s1,goal)
}