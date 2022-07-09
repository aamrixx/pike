package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("mkdir : 1 operand required\n")
		os.Exit(1)
	}

	if err := os.Mkdir(os.Args[1], os.ModePerm); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	fmt.Printf("created directory '%s'\n", os.Args[1])
}
