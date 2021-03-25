package main

import (
	"fmt"
	"os"
)

func Command(args []string) {
	for _, val := range os.Environ() {
		fmt.Printf("%s\n", val)
	}
	fmt.Printf("Args: %v\n", args)
}
