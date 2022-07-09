// This ls command needs more work
// but so far good
// TODO : Hide dotfiles

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ls : 1 operand required.\n")
		return
	}

	f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
        return
    }

    files, err := f.Readdir(0)
    if err != nil {
        log.Fatal(err)
        return
    }

    for _, v := range files {
		if v.IsDir() {
			fmt.Printf("\033[1;34m%s\033[0m\n", v.Name())
		} else {
			fmt.Printf("%s\n", v.Name())
		}
    }
}
