// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	fmt.Println(forEachNode("toc", doc, startElement, endElement))
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(id string, n *html.Node, pre, post func(n *html.Node, id string) bool) *html.Node {
	
	if pre != nil {
		if pre(n, id) {
			result := n
			return result
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		x := forEachNode(id, c, pre, post)
		if x != nil {
			return x
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

//!-forEachNode

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(id, doc, startElement, endElement)
}


//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && len(n.Attr) > 0 {
		fmt.Println("Lets try this: " + n.Data)
		for i:=0; i<len(n.Attr); i++ {
			fmt.Println(n.Attr[i].Key == "id")
			fmt.Println(n.Attr[i].Val == id)
			if n.Attr[i].Key == "id" && n.Attr[i].Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && len(n.Attr) > 0 {
		for i:=0; i<len(n.Attr); i++ {
			if n.Attr[i].Key == "id" && n.Attr[i].Val == id {
				return true
			}
		}
	}
	return false
}

//!-startend
