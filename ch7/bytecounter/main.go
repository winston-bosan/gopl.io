// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"fmt"
	"io"
)

//!+bytecounter

type ByteCounter int
type WordCounter int
type LineCounter int
type sFunc func([]byte, bool) (int, []byte, error)

type CountCounter struct {
	w 				io.Writer
	inner_counter 	int64
}

func countScan(p []byte, scanFunc sFunc) int {
	counter := 0
	holder := true
	for holder {
		idx_advance, token, _ := scanFunc(p, true)
		if len(token) == 0 {
			holder = false
			break
		}
		p = p[idx_advance:]
		counter++
	}
	return counter
}

func (c *CountCounter) Write(p []byte) (int, error) {
	c.inner_counter += int64(len(p))
	c.w.Write(p)
	return len(p), nil
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	num_lines := countScan(p, bufio.ScanLines) // convert int to ByteCounter
	*c += LineCounter(num_lines)
	return num_lines, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	num_words := countScan(p, bufio.ScanWords) // convert int to ByteCounter
	*c += WordCounter(num_words)
	return num_words, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	temp_writer := &CountCounter{w, 0} 
	return temp_writer, &temp_writer.inner_counter
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	var d LineCounter
	d.Write([]byte("HELLO.\n is this how itworks? \n nope."))
	fmt.Println(d) // "3"

	var e WordCounter
	e.Write([]byte("hellos mein germany frued"))
	fmt.Println(e) // "3"

	z, count := CountingWriter(&e)
	fmt.Println(*count) //why the pointer???
	z.Write([]byte("MEH MEH MEH"))
	fmt.Println(*count)
}
