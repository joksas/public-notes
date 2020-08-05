package main

import (
	"fmt"
	"os"
	"popcount"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		number, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		popCount := popcount.PopCount(number)
		fmt.Printf("Number: %d, pop-count: %d\n", number, popCount)
	}
}
