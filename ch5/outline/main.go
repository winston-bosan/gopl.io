// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	stack_stack := make([]string, 10)
	outline(&stack_stack, nil, doc)
	el_mapping(stack_stack)
}

func outline(stack_stack *[]string, stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		*stack_stack = append(*stack_stack, stack[len(stack)-1])
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack_stack, stack, c)
	}
}

func el_mapping(stack_stack []string) {
	html_map := map[string]int{}
	// @. Loop through the stack_stack
	for i := 0; i < len(stack_stack); i++ {
		// 1. IF key, ADD-INPLACE to map:key
		html_map[stack_stack[i]]++
	}
	// fmt.Println(html_map)
}
//!-
