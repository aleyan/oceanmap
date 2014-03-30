package main

import "fmt"
import "os"

func scrapeEverything() {
	url := "http://digitalocean.com"
	initial := new(Page)
	initial.Url = url
	unvisited = append(unvisited, initial)
	TraverseGraph()

	fmt.Println()
	fmt.Printf("Found out about %v pages with %v left\n", len(pages), len(unvisited))
	PrintGraph()
	fmt.Println("that is all folks!")
}

func scrapeOnePage(url string) {
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
