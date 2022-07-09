package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("date : no operands required\n")
		return
	}

	currentTime := time.Now()

	fmt.Printf("%s\n", currentTime.Format("Mon  02  Jan  15:04:05  BST  2006"))
}
