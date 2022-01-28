// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
// Exercise 1.4: Modified dup2 to print the names of all the files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%s\n%s\n\n", line, files)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func countLines(f *os.File, counts map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !contains(counts[input.Text()], filename) {
			counts[input.Text()] = append(counts[input.Text()], filename)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
