package main

import (
	"strconv"
	"syscall"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("chmod : expected 2 operands (file and mode)\n")
		os.Exit(1)
	}

	mode, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	err = syscall.Chmod(os.Args[1], uint32(mode))
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s -> %s\n", os.Args[1], os.Args[2])
}
