package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	month := []*github.Issue{}
	year := []*github.Issue{}
	overyear := []*github.Issue{}

	for _, item := range result.Items {

		if time.Since(item.CreatedAt).Hours() < 720.0 {
			month = append(month, item)
		}

		if time.Since(item.CreatedAt).Hours() < 720.0*12 {
			year = append(year, item)

		}

		if time.Since(item.CreatedAt).Hours() >= 720.0*12 {
			overyear = append(overyear, item)
		}
	}

	fmt.Printf(" ↓↓↓↓ Less than one month  ↓↓↓↓ \n")

	for _, item := range month {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf(" ↓↓↓↓ Less than one year ↓↓↓↓  \n")

	for _, item := range year {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf(" ↓↓↓↓ Over one year ↓↓↓↓ \n")

	for _, item := range overyear {

		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}
