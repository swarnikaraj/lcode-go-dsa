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

func isAnagram(s string, t string) bool {
    // Length of both strings should be the same for them to be anagrams
    if len(s) != len(t) {
        return false
    }
    
    // Create a map to count occurrences of characters in s
    charCount := make(map[rune]int)
    for _, char := range s {
        charCount[char]++
    }
    
    // Decrement counts for characters in t
    for _, char := range t {
        charCount[char]--
        // If a character count becomes negative, strings are not anagrams
        if charCount[char] < 0 {
            return false
        }
    }
    
    // Check if all counts are zero, indicating anagrams
    for _, count := range charCount {
        if count != 0 {
            return false
        }
    }
    
    return true
}

func main() {
	s := "a good   example"
	reverseWords(s)
s1:= "abcde"
 goal := "cdeab"
	rotateString(s1,goal)
}