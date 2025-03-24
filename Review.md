# Review for futher modifications to be made to the code. 

1. Make use of an env file for the baseURL and the API
2. Merge the fields here
    ```go 
    type data struct {
            Amount             string `json:"amount"`
            Currency           string `json:"currency"`
            From               string `json:"from"`
            Description        string `json:"description"`
            External_reference string `json:"external_reference"`
        }
    ```
        and here 
    ```go
        // put all of this in a globla form and merge it with that of 48 - 53
        var output struct {
            Reference          string `json:"reference"`
            Status             string `json:"status"`
            Amount             int    `json:"amount"`
            Currency           string `json:"currency"`
            Operator           string `json:"operator"`
            Code               string `json:"code"`
            Operator_reference string `json:"operator_reference"`
        }
    ````
3. Make use of Modules or packages to simplify your main file that runs the code and calls the othere API file. 
Read on this https://stackoverflow.com/questions/29898400/import-struct-from-another-package-and-file-golang 

4. Move line 55 -- 81 to a new module or file so as to make the codebase simpler and more readable. 

## Resources
- https://go.dev/doc/tutorial/create-module
- https://gobyexample.com/ 
- https://go.dev/ref/spec#Packages