package main

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input:
var strs = []string{"flower","flow","flight"}
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.


func longestCommonPrefix(strs []string) string {
    var res string

    if len(strs) == 0 {
        return res
    }

    for j := 0; j < len(strs[0]); j++ {
        for i := 1; i < len(strs); i++ {
            
            if j >= len(strs[i]) || strs[i][j] != strs[0][j] {
                return res
            }
        }
        res = res + string(strs[0][j])
    }

    return res
}


func main() {
longestCommonPrefix(strs)
}