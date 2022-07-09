package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("rm : 1 operand required\n")
		return
	}

	if err := os.RemoveAll(os.Args[1]); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("removed '%s'\n", os.Args[1])
}
