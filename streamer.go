package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(getContent("07112"))
}

func getContent(zipcode string) string {

	var result map[string]interface{}
	json.Unmarshal([]byte(call(zipcode)), &result)
	return "<div>" + strconv.FormatFloat(result["current"].(map[string]interface{})["temp_c"].(float64), 'f', 2, 64) + "</div>"
}

func call(zipcode string) []byte {
	resp, _ := http.Get("http://api.weatherapi.com/v1/forecast.json?key=2b23e7071c304d53b82160733203007&q=" + zipcode + "&day=1")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}
