package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*
This function makes a payment request to the compay api

It takes 2 struct arguments all of type Data (defined in the data.go file)
 1. the requestData struct which holds payment details like the amount request, currency, and description
 2. the result struct which holds the response of the payment request i.e the reference operator etc

It returns the reference of the transaction which is then used to check the transaction status
*/

// This function initiates a payment request
func RequestPayment(requestData, result Data) string {

	//load the api-key from the .env file, handle errors if they occur
	key, baseUrl, err := LoadCredentials()
	if err != nil {
		fmt.Println("failed to load apikey or baseUrl", err)
		os.Exit(1)
	}
	//convert the payment details into json
	//because the body of the request expects json data, catch any error that
	//may occur
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Println(string(jsonData))
	}

	//initialize an http client.
	// http.Client{} is a struct in Go's net/http package
	client := &http.Client{}

	//create an http request using http.NewRequest function from net/http

	url := fmt.Sprintf("%vcollect/", baseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("error occured:", err)
	}

	//adding custom headers to http.newrequest()
	//authorization and content type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", key)

	//send an http request and return the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error occured", err)
	}
	defer resp.Body.Close()

	// decode the body of the response to a readable format using the sresult tructs
	json.NewDecoder(resp.Body).Decode(&result)

	//obtain the transaction reference from the response body
	//this reference will be used to check the status of the transaction
	return result.Reference

}
