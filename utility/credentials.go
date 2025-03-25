package utility

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
This go file loads the credentials needed to make http requests to the campay api
*/

// This function loads the api key and the base url from the .env file
func LoadCredentials() (string, string, error) {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return "", "", fmt.Errorf("failed to load env file: %w", err)
		}
	}

	key := os.Getenv("key")

	baseUrl := os.Getenv("baseUrl")

	if key == "" {
		return "", "", fmt.Errorf("API_KEY is not set")
	}

	if baseUrl == "" {
		return "", "", fmt.Errorf("BaseUrl not found")
	}

	return key, baseUrl, nil
}
