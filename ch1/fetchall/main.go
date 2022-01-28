// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
// Exercise 1.10: Modified to print its output to a file.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(i, url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(i int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	resp.Body.Close() // don't leak resources
	err = os.WriteFile(strings.Split(url, "://")[1] + ".html", b, 0777)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %s", secs, url)
}

//!-
