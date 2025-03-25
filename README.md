# APIs-with-Go
This project uses Golang to make Http Requests to the various Campay endpoints

# Project structure and functioning

The project is contained in a go module with folder named utility which has all the go files defined in the utility package.

In the utility folder we have 4 go files:
 1. credentials.go: used to load the api key and base url form the .env file
 2. data.go: which holds a public struct to hold the input data needed to make a payment request. this struct is also used accross the program to hold the response of each htpp request
 3. request_payment.go file: used to make a payment request
 4. check_transaction_status.go file: which checks the transaction status

 The main.go file calls other functions from the utility package to successfully make a payment request and check the transaction reference. 
