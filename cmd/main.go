package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Order struct {
	CustomerDetails struct {
		CustomerID    string `json:"customer_id"`
		CustomerPhone string `json:"customer_phone"`
	} `json:"customer_details"`
	OrderCurrency string  `json:"order_currency"`
	OrderAmount   float64 `json:"order_amount"`
}
type Customer struct {
	CustomerPhone string `json:"customer_phone"`
	CustomerEmail string `json:"customer_email"`
	CustomerName  string `json:"customer_name"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	clientId := os.Getenv("CLIENT_ID")
	fmt.Println(clientId)
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)

	customer := Customer{
		CustomerPhone: "9999999999",
		CustomerEmail: "prasun@example.com",
		CustomerName:  "Prasun",
	}

	jsonData, err := json.Marshal(customer)
	fmt.Println(jsonData)
	if err != nil {
		fmt.Println("Error marshalling customer data:", err)
		return
	}

	url := "https://sandbox.cashfree.com/pg/customers"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-version", "2023-08-01")
	req.Header.Set("X-Client-Secret", secretKey)
	req.Header.Set("X-Client-Id", clientId)

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

	// Print the response
	fmt.Println("Response:", string(body))
}
