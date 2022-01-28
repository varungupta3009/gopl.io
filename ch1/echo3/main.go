// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
// Exercise 1.3: Modified to measure time efficiency b/w multiple algorithms.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()
	result := strings.Join(os.Args[1:], " ")
	fmt.Println(result)
	fmt.Printf("Time taken for string.Join(): %dμs\n\n", time.Since(start).Microseconds())
	start = time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	result = s
	fmt.Println(result)
	fmt.Printf("Time taken for iterative algorithm: %dμs\n", time.Since(start).Microseconds())
}

//!-
