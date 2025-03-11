package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// This function loads the api key from the .env file
func run() (string, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return "", fmt.Errorf("failed to load env file: %w", err)
		}
	}

	key := os.Getenv("key")

	if key == "" {
		return "", fmt.Errorf("API_KEY is not set")
	}

	return key, nil
}

// This function initiates a payment request
func requestPayment() (string, string) {

	//url to request payment using campay
	baseurl := "https://demo.campay.net/api/collect/"

	//load the api-key from the .env file, handle errors if they occur
	key, err := run()
	if err != nil {
		fmt.Println("failed to load apikey", err)
		os.Exit(1)
	}

	// get the required data needed to make a payment request.
	// json struct tags (e.g `json:"name"`) are used to allow better control
	// of field names, formatting when dealing with json
	type data struct {
		Amount             string `json:"amount"`
		Currency           string `json:"currency"`
		From               string `json:"from"`
		Description        string `json:"description"`
		External_reference string `json:"external_reference"`
	}

	//ask users to enter payment details
	var amount string
	var number string
	var description string
	var reference string

	currency := "XAF"

	fmt.Println("enter amount")
	fmt.Scanln(&amount)

	fmt.Println("enter number")
	fmt.Scanln(&number)

	fmt.Println("enter description")
	fmt.Scanln(&description)

	fmt.Println("enter reference")
	fmt.Scanln(&reference)

	//create an instance of the struct with the payment details
	requestData := data{
		Amount:             amount,
		Currency:           currency,
		From:               number,
		Description:        description,
		External_reference: reference,
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
	req, err := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonData))

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

	//create a struct to hold the response body of the http request
	var result struct {
		Reference string `json:"reference"`
		Ussd_code string `json:"ussd_code"`
		Operator  string `json:"operator"`
	}

	// decode the body of the response to a readable format using the structs
	json.NewDecoder(resp.Body).Decode(&result)

	//obtain the transaction reference form the response body
	//this reference will be used to check the status of the transaction
	return result.Reference, key

}

func checkTransactionStatus() {

	//get the transaction reference
	transactionReference, key := requestPayment()

	//now check the transaction staatus
	statusUrl := fmt.Sprintf("https://demo.campay.net/api/transaction/%v", transactionReference)

	request, err := http.NewRequest("GET", statusUrl, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", key)
	client2 := &http.Client{}

	response, err := client2.Do(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	var output struct {
		Reference          string `json:"reference"`
		Status             string `json:"status"`
		Amount             int    `json:"amount"`
		Currency           string `json:"currency"`
		Operator           string `json:"operator"`
		Code               string `json:"code"`
		Operator_reference string `json:"operator_reference"`
	}
	json.NewDecoder(response.Body).Decode(&output)
	transactionStatus := output.Status

	time.Sleep(30 * time.Second)
	fmt.Println("transaction Status", transactionStatus)
}

func main() {

	checkTransactionStatus()
}
