package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Printf("echo : 1 operand expected")
		fmt.Printf("echo [TEXT] [>>] [FILE}")
		return
	}

	if len(os.Args) == 2 {
		fmt.Printf("%s\n", os.Args[1])
		return
	}

	if len(os.Args) == 4 {
		f, err := os.ReadFile(os.Args[2])
    	if err != nil {
       		log.Fatal(err)
			return
		}

		err = os.WriteFile(os.Args[1], f, 0644)
    	if err != nil {
        	log.Fatal(err)
			return
    	}
	}
}
