package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func GetMagnetLinks(httpBody io.Reader, limit int) []string {
	links := make([]string, 0)
	elem := html.NewTokenizer(httpBody)

	for {
		tokenType := elem.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := elem.Token()
		if token.DataAtom.String() == "a" && len(links) < limit {
			for _, attr := range token.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "magnet") {
					links = append(links, attr.Val)
					elem.Next()
				}
			}
		}
	}
}
