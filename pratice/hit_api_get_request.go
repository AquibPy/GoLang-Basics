package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://redis-4eqp.onrender.com/entries"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers to the request
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJtb2hkYXF1aWIzMUBnbWFpbC5jb20iLCJleHAiOjE3MTU5Njc4Njh9.Wok3qkMKL9ELH6BPmgfwqlHB2BmaaxcM-6Qc1YYAPTg")

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Response:", string(body))
}
