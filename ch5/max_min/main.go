package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

//!+
func main() {
	list_nums := []int{1,23,4,5,6,7,435,34,534,54,-34892384392}
	list_str := []string{"-MEH-", "fdfsa", "fodod", "923jdi"}

	fmt.Println(max(list_nums...))
	fmt.Println(min(list_nums...))
	fmt.Println(join("@@@", list_str...))

	resp, err := http.Get("https://go.dev")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Getting https://go.dev ......")

	list_tags := []string{"img", "h1", "h2", "h3"}
	wanted_nodes := ElementsByTagName(doc, list_tags...)
	for i := 0; i < len(wanted_nodes); i++ {
		fmt.Println(wanted_nodes[i])
	}
	

}

func ElementsByTagName(doc *html.Node, names ...string) (results []*html.Node) {
	for i := 0; i < len(names); i++ {
		fmt.Println("Now seraching for: " + names[i])
		go_through_doc(&results, names[i], doc)
	}
	return results
}

func go_through_doc(return_bucket *[]*html.Node, target string, n *html.Node) {
	if n.Data == target {
		fmt.Println(n.Data)
		*return_bucket = append(*return_bucket, n) // push tag
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		go_through_doc(return_bucket, target, c)
	}
}



func join(joinery string, strings ...string) string {
	shit_holder := ""
	for i := 0; i < len(strings); i++ {
		shit_holder = shit_holder + joinery + strings[i]
	}
	return shit_holder
}

func max(nums ...int) int {
	currentMax := nums[0]

	for i := 0; i<len(nums); i++ {
		if nums[i] > currentMax{
			currentMax = nums[i]
		}
	}

	return currentMax
}

func min(nums ...int) int {
	currentMax := nums[0]

	for i := 0; i<len(nums); i++ {
		if nums[i] < currentMax{
			currentMax = nums[i]
		}
	}

	return currentMax
}