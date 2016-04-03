package main

import (
	"strings"
)

var defaultUrl = "https://pirateproxy.tv"

type SearchParams struct {
	search string
	url    string
}

func (s *SearchParams) Create(params []string) {
	s.search = params[0]

	if len(params) > 1 {
		s.url = createValidUlr(params[1])
	} else {
		s.url = defaultUrl
	}
}

func fillInWhiteSpaces(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func createValidUlr(urlStr string) string {

	prefix := "http://"
	if !strings.HasPrefix(urlStr, prefix) {
		urlStr = prefix + urlStr
	}

	suffix := "/"
	if strings.HasSuffix(urlStr, suffix) {
		urlStr = urlStr[:len(urlStr)-len(suffix)]
	}

	return urlStr
}

func (s *SearchParams) ConstructUrl() string {
	return s.url + "/search/" + fillInWhiteSpaces(s.search)
}
