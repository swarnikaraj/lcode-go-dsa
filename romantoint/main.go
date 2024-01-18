package main

import "fmt"

var rommap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
func romanToInt(s string) int {

	var rommap = map[byte]int{

		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	
	for i:=0;i<len(s);i++{
         
		 
          if i+1 < len(s) && rommap[byte(s[i])] < rommap[byte(s[i+1])]{
			tempres:= rommap[byte(s[i+1])]-rommap[byte(s[i])]
			res=res+ tempres
			i++
		  }else{
          res = res + rommap[byte(s[i])]
		  }
		
	}
	fmt.Print(res)
	return res

}

func main() {
	romanToInt("MCMXCIV")
}