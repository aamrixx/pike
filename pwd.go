package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("pwd : no operands required\n")
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s\n", dir)
}
