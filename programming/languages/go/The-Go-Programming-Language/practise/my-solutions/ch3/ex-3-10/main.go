// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma_buffer(os.Args[i]))
	}
}

func comma_buffer(s string) string {
	n := len(s)
	numSegments := n / 3
	prefix := n % 3
	var buf bytes.Buffer
	buf.WriteString(s[:prefix])
	for i := 0; i < numSegments; i++ {
		if prefix != 0 || i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[prefix+i*3 : prefix+(i+1)*3])
	}

	return buf.String()
}
