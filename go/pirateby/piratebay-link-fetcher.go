package main

import (
	"fmt"
	"net/http"
)

func printResults(links []string) {
	if len(links) == 0 {
		fmt.Println("No results! Sorry :(")
	} else {
		printLinks(links)
	}
}

func printLinks(links []string) {
	fmt.Println("Here are links you are looking for. Enjoy!\n")

	for _, link := range links {
		fmt.Println(link, "\n")
	}
}

func main() {

	params := SearchParams{resultLimit: limit, host: defaultHost}
	url := params.ConstructUrl("originals s01e02")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	links := GetMagnetLinks(resp.Body, limit)
	printResults(links)

}
