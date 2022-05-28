package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var baseUrl = "https://google-finance4.p.rapidapi.com"

func Greet(text string) string {
	return text + " \nCall /search?q={add any company} to get a financial report"
}

func token() string {
	return os.Getenv("TOKEN")
}

func setUpClient(url string, method string) (*http.Client, *http.Request) {
	// Create a new request using http
	req, _ := http.NewRequest(method, baseUrl+url, nil)
	log.Print("==========>", baseUrl+url)

	// add authorization header to the req

	req.Header.Add("X-RapidAPI-Host", "google-finance4.p.rapidapi.com")

	req.Header.Add("X-RapidAPI-Key", token())

	req.Header.Add("Accept", "application/json")
	// Send req using http Client
	client := &http.Client{}

	return client, req
}
func Get(url string) ([]byte, error) {

	// response, err := http.Get(url)

	client, req := setUpClient(url, "GET")

	response, err := client.Do(req)

	if err != nil {
		print(err)
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)

}
