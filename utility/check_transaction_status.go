package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

/*
This go file defines the checktransaction status function

The function takes 2 arguments
	1. the reference of the transacction whose status is to be checked
	2. A struct which will hold the response body

The function returns the transaction status of the specified trasaction
*/

func CheckTransactionStatus(reference string, output Data) {

	//load the api-key from the .env file, handle errors if they occur
	key, baseUrl, err := LoadCredentials()
	if err != nil {
		fmt.Println("failed to load apikey or baseUrl", err)
		os.Exit(1)
	}

	//now check the transaction staatus

	url := fmt.Sprintf("%vtransaction/", baseUrl)
	statusUrl := fmt.Sprintf("%v%v", url, reference)

	request, err := http.NewRequest("GET", statusUrl, nil)
	if err != nil { // put also in .env file
		fmt.Println("error:", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", key)
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	json.NewDecoder(response.Body).Decode(&output)
	transactionStatus := output.Status

	time.Sleep(60 * time.Second)
	fmt.Println("transaction Status", transactionStatus)
}
