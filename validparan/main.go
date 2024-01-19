package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input:
var s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false
func isValid(s string) bool {
    var stack []string

    if len(s)%2 != 0 {
        return false
    }

    pairs := map[string]string{
        "(": ")",
        "[": "]",
        "{": "}",
    }

    for i := 0; i < len(s); i++ {
        char := string(s[i])

        if _, isOpen := pairs[char]; isOpen {
            stack = append(stack, char)
        } else {
           
            if len(stack) == 0 || pairs[stack[len(stack)-1]] != char {
                return false
            }

           
            stack = stack[:len(stack)-1]
        }
    }

 
    return len(stack) == 0
}

func main(){
isValid(s)
}