package main

import (
	"strings"
)

var defaultHost = "https://pirateproxy.tv"
var limit = 5

type SearchParams struct {
	search      string
	resultLimit int
	host        string
}

func fillInWhiteSpaces(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func (s *SearchParams) ConstructUrl(search string) string {
	return s.host + "/search/" + fillInWhiteSpaces(search)
}
