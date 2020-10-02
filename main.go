package main

import (
	"fmt"
	"os"
)

func main() {
	// Fictional secret
	apiKey := os.Getenv("API_KEY")
	fmt.Println(apiKey)
}
