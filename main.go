package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Product struct to hold product data
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// in-memory store simulation
// TODO: use a map or slice to store the product data

// helper function to apply discount
func applyDiscount(price float64, percentage float64) float64 {
	// For simplicity, just subtract the discount
	// In real scenarios, edge cases should be considered
	return price - (price * percentage / 100)
}

func addProduct(product *Product) int {
	// TODO: implement the addProduct function based on requirements
	// - Check for duplicates
	// - Validate product fields
	// - Add to in-memory store

	return 0 // success
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	// parse the request body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// add product
	result := addProduct(&product)
	if result == -1 {
		http.Error(w, "Product already exists or invalid input", http.StatusBadRequest)
		return
	}

	// respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product added successfully"))
}

func main() {
	http.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			addProductHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
