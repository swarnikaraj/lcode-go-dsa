package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
) 
func encrypt(n int) int{
    length:=0
    max:=math.MinInt
    for n>0{
        lastdigit:=n%10
		fmt.Println(lastdigit>max,"lastdigit")
        if max<lastdigit{
            max=lastdigit
        }
        length++
        n=n/10
    }
   
     numstr:=""
    for i:=0;i<length;i++{
        numstr=numstr + strconv.Itoa(max)
    }
    
    intans,_:=strconv.Atoi(numstr)
    return intans
    
}
  
func sumOfEncryptedInt(nums []int) int {
    sum:=0
    for i:=0;i<len(nums);i++{
        en:=encrypt(nums[i])
		fmt.Println(en)
        sum=sum + en
    }
  
    return sum
}


func minimizeStringValue(s string) string {
  
    ans:=""
    freqMap:=make(map[string]int)
    minFrequency := math.MaxInt64
	var minAlpha string
    var found bool=false
    i:=0
    
    for i<len(s){
        if s[i]=='?'{
            
           for ch := 'a'; ch <= 'z'; ch++ {
      
            if _, ok := freqMap[string(ch)]; !ok {
				
             minAlpha=string(ch)
			
             found=true
             break
             }
            }

	

            if !found{
                
            
	     for alpha, freq := range freqMap {
		     if freq < minFrequency {
			   minFrequency = freq
			    minAlpha = alpha
		    }
	          }
            }
			
            ans=ans+string(minAlpha)
			freqMap[minAlpha]=1
       
        }else{
            ans=ans+string(s[i])
            if mapped, ok:= freqMap[string(s[i])];ok{
               fmt.Println(mapped," mapped is this")
                freqMap[string(s[i])]=mapped+1
            }else{
                freqMap[string(s[i])]=1
            }
        }
        
        i++
    }
   
    return ans
}

func isPalin(str string) bool{
    i:=0;
    j:=len(str)-1

    for i<=j{
        if str[i]!=str[j]{
            return false
        }
        
        i++
         j--
    }
    
    
    return true
}

// 100248. Existence of a Substring in a String and Its Reverse
// User Accepted:4711
// User Tried:5359
// Total Accepted:4739
// Total Submissions:6142
// Difficulty:Easy
// Given a string s, find any substring of length 2 which is also present in the reverse of s.

// Return true if such a substring exists, and false otherwise.

 

// Example 1:

// Input: s = "leetcode"

// Output: true

// Explanation: Substring "ee" is of length 2 which is also present in reverse(s) == "edocteel".

// Example 2:

// Input: s = "abcba"

// Output: true

// Explanation: All of the substrings of length 2 "ab", "bc", "cb", "ba" are also present in reverse(s) == "abcba".

// Example 3:

// Input: s = "abcd"

// Output: false

// Explanation: There is no substring of length 2 in s, which is also present in the reverse of s.

func isSubstringPresent(s string) bool {
    if len(s)<2{
        return false
    }
    reversestr:=""
  
    for j:=len(s)-1;j>=0;j--{
        reversestr=reversestr+string(s[j])
    }
   
    
    for i := 0; i <= len(s)-2; i++ {
        substrings1:= s[i:i+2]
        if strings.Contains(reversestr,substrings1){
            return true
        }
	}
    
   
    
    
    return false
    
}

func countSubstrings(s string, c byte) int64 {
   var count int64=0
    for i:=0;i<len(s);i++{
        for j:=i+1;j<=len(s);j++{
            sbstr:=s[i:j]
            if sbstr[0]==sbstr[len(sbstr)-1] &&  sbstr[len(sbstr)-1]==c{
                count++
            }
        }
    }
    
    return count
    
}

func countSubstrings2(s string, c byte) int64 {
  
    var count int64

    // Count occurrences of the character 'c' in the string
    var occurrences int64
    for i := 0; i < len(s); i++ {
        if s[i] == c {
            occurrences++
        }
    }

    // Calculate the number of substrings starting and ending with 'c'
    count = occurrences * (occurrences + 1) / 2

    return count
    
}
    
func minimumDeletions(word string, k int) int {
    output:=0
    obj:=make(map[byte]int)
    obj2:=make(map[byte]int)

    for i:=0;i<len(word);i++{
        if mapped, ok:=obj[word[i]];ok{
            obj[word[i]]=mapped +1
        }else{
               obj[word[i]]=1
        }
    }
    
    for i:=0;i<len(word);i++{
		target:=word[i]
		for j:=0;j<len(word);j++{
			if i!=j{
              mappedT, _:=obj[target]
			  mappedJ, _:=obj[word[j]]
			  if int(math.Abs(float64(mappedT)-float64(mappedJ)))>k{
                 obj2[word[i]]=1
			  }
			}
		}
	}
    
    
    for num :=range obj2{
		fmt.Println(num)
        output++
    }
    return output
}

func main(){
// nums := []int{1,2,3}
// sumOfEncryptedInt(nums)
// input:="???"
// sol:=minimizeStringValue(input)
// fmt.Println(sol)
// st:="abcd"
// isSubstringPresent(st)

word := "dabdcbdcdcd"
k := 2
minimumDeletions(word, k)
}