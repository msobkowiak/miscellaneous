package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Url struct {
	url, city, countryCode string
}

func buildUrl(url Url) string {
	str := url.url + url.city + "," + url.countryCode + "&units=metric"
	return str
}

func main() {

	// constract api url
	urlParams := Url{"http://api.openweathermap.org/data/2.5/weather?q=", "porto", "pt"}
	url := buildUrl(urlParams)

	// get a data from weather api
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	jsonStr := string(body[:])

	// parse json
	var weather map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &weather)
	temp := weather["main"].(map[string]interface{})["temp"]
	fmt.Println("The temperature at", urlParams.city, "is", temp, "ºC")
}
