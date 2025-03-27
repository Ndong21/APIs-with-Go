package utility

/*
This go file defines a public struct.

This struc will be used to;
 1. store the data needed to make a payment request collected from the users like amount, currency, phone number, description and external reference
 2. hold the results obtained by making the payment request like the transaction reference, ussd code and operator
 3. hold the output obtained when checking the transaction reference like amount, currency, transaction status etc

The struct uses json struct tags which allow better control of fieldnames when turning the data to json to send in an http request

This struct will be used in several scenarios, so to avoid some errors that might occur when some properties of the struct are not used, i made all the properties optional by using 'omitempty'

This struct is therefore imported into other files in the utility package as well as main
*/

type Data struct {
	Amount             string `json:"amount,omitempty"`
	Currency           string `json:"currency,omitempty"`
	From               string `json:"from,omitempty"`
	Description        string `json:"description,omitempty"`
	External_reference string `json:"external_reference,omitempty"`
	Reference          string `json:"reference,omitempty"`
	Status             string `json:"status,omitempty"`
	Operator           string `json:"operator,omitempty"`
	Code               string `json:"code,omitempty"`
	Operator_reference string `json:"operator_reference,omitempty"`
}
