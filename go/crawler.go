package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func Avarage(nums []float64) float64 {
	total := Sum(nums)
	return total / float64(len(nums))
}

func Sum(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func GetPrices(httpBody io.Reader) []string {
	prices := make([]string, 0)
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return prices
		}
		token := page.Token()
		if token.DataAtom.String() == "strong" {
			for _, attr := range token.Attr {
				if attr.Key == "class" && attr.Val == "value" {
					page.Next()
					prices = append(prices, page.Token().String())
				}
			}
		}
	}
}

func skipCurrencySymbol(price string) string {
	return strings.TrimSuffix(price, "€")
}

func converStringToFloat(strs []string) []float64 {
	numbers := make([]float64, 0)
	for _, str := range strs {
		tmp := skipCurrencySymbol(str)
		num, _ := strconv.ParseFloat(tmp, 64)
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	resp, err := http.Get("https://www.uniplaces.com/accommodation/lisbon")
	if err != nil {
		panic(err.Error())
	}

	prices := GetPrices(resp.Body)
	fmt.Println("All prices: ", prices)
	defer resp.Body.Close()

	numbers := converStringToFloat(prices)
	fmt.Println("Sum of prices: ", Sum(numbers), "€")
	fmt.Println("Avagare price: ", Avarage(numbers), "€")
}
