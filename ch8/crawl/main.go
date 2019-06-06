package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl.io/ch5/links"
)

type Link struct {
	Url   string
	Depth int
}

func crawl(link Link) []Link {
	fmt.Println(link)
	list, err := links.Extract(link.Url)
	if err != nil {
		log.Print(err)
	}

	var depth = link.Depth + 1
	var result []Link
	for _, url := range list {

		result = append(result, Link{url, depth})
	}
	return result
}

func main() {
	worklist := make(chan []Link)
	unseenLinks := make(chan Link)
	var workingLinks []Link
	go func() {

		for _, url := range os.Args[1:] {
			workingLinks = append(workingLinks, Link{url, 1})
		}

		worklist <- workingLinks

	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.Depth <= 3 {
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}

			}
		}()
	}

	seen := make(map[Link]bool)
	for list := range worklist {

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
