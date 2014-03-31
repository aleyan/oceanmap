package main

import (
	"fmt"
	"os"
	"strings"
)

func scrapeEverything() {
	url := "http://digitalocean.com"
	initial := new(Page)
	initial.Url = url
	unvisited = append(unvisited, initial)
	TraverseGraph()

	fmt.Println()
	fmt.Printf("Found out about %v pages and scanned %v of them\n", len(pages), len(pages)-len(unvisited)+1)
	PrintGraph()
	fmt.Println("that is all folks!")
}

func scrapeOnePage(url string) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	fmt.Println(url)
	initial := new(Page)
	initial.Url = url
	TraversePage(initial)
	PrintPageLinks(initial)
}

func main() {
	if len(os.Args) == 2 {
		scrapeOnePage(os.Args[1])
	} else {
		scrapeEverything()
	}

}
