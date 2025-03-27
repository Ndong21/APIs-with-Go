package main

import (
	"fmt"

	//import the package "utility"
	. "example.com/api/utility" // using . infront helps us omit the package name when calling the exported identifiers(fxns and variables)
)

func main() {

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
	requestData := Data{
		Amount:             amount,
		Currency:           currency,
		From:               number,
		Description:        description,
		External_reference: reference,
	}

	//create an instance of the Data struct
	var result Data //result will be used to store response boy from the requestpayment function
	var output Data // output will be used to store the response from the checktransactionstatus function

	//get the transaction reference from the requestpayment function
	transactionReference := RequestPayment(requestData, result)

	fmt.Println("Payment request successfully initiated")
	fmt.Println("Now checking the status of the transaction")
	fmt.Println(".........")

	//check the transaction status of the operation
	CheckTransactionStatus(transactionReference, output)
}
