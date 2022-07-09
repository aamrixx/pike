// Currently this implementation
// Does change perms but does it under
// root I think.

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
		return
	}

	mode, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = syscall.Chmod(os.Args[1], uint32(mode))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s -> %s\n", os.Args[1], os.Args[2])
}
