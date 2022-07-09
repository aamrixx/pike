package main

import (
	"fmt"
	"log"
	"strconv"
	"syscall"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("kill : 1 operand required\n")
		fmt.Printf("kill [PID] [SIGNAL]\n")
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
    if err != nil {
    	log.Fatal(err)
		return
	}

	var sig syscall.Signal = 1

	if err := syscall.Kill(pid, sig); err != nil{
		log.Fatal(err)
		return
	}

	fmt.Printf("killed process with id %s\n", os.Args[1])
}
