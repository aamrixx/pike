package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("pwd : no operands required\n")
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s\n", dir)
}
