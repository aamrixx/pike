package main

import (
	"fmt"
	"os"
)

func reverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    return string(rns)
}

func main() {
	if len(os.Args) > 3 || len(os.Args) == 1 {
		fmt.Printf("basename : 1 operand required\n")
		fmt.Printf("basename [SUFFIX] [STRING]\n")
		return
	}

	var substr string
	chars := []rune(os.Args[1])
    var start int

	if len(os.Args) == 2 {
		start = len(chars)-1
	} else if len(os.Args) == 3 {
		chars = []rune(os.Args[2])
		start = len(os.Args[2])-len(os.Args[1])-1
	}

	for i := start; i > 0; i-- {
    	if string(chars[i]) == "/" {
			break
		}

	    substr += string(chars[i])
    }

	fmt.Printf("%s\n", reverse(substr))
}
