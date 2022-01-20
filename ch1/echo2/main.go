// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
// Exercise 1.2: Modified to print index along with argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, arg := range os.Args[1:] {
		fmt.Println(ind, arg)
	}
}

//!-
