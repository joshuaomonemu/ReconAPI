package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	resp, err := http.Get("https://api.paystack.co/bank?Authorization=Bearer YOUR_SECRET_KEY")

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
}
