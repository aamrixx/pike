package main

import (
	"log"
	"os"
	"io"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("cp : 2 operands expected\n")
		fmt.Printf("cp [SOURCE] [DESTINATION]\n")
		return
	}

    buf := make([]byte, 1024)

    fin, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    defer fin.Close()

    fout, err := os.Create(os.Args[2])
    if err != nil {
        log.Fatal(err)
    }

    defer fout.Close()

    for {

        n, err := fin.Read(buf)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }

        if n == 0 {
            break
        }

        if _, err := fout.Write(buf[:n]); err != nil {
            log.Fatal(err)
        }
   	}

	fmt.Printf("copied '%s' to '%s'\n", os.Args[1], os.Args[2])
}
