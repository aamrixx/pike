package main

import (
	"fmt"
	"log"
	"os/user"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("whoami : no operands required\n")
		return
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s\n", currentUser.Username)
}
