package main

import (
	"fmt"
	"os"

	"github.com/hayajo/which"
)

func main() {
	var exit int

	if len(os.Args) < 2 {
		exit = 1
		os.Exit(exit)
	}

	for _, p := range os.Args[1:] {
		path, err := which.Which(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			exit = 1
			continue
		}
		fmt.Println(path)
	}

	if exit != 0 {
		os.Exit(exit)
	}
}
