package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Harga int `json:"harga"`
	Stock int `json:"stock"`
}

var product = []Product{
	{ID: 1, Name: "indomie", Harga: 5000, Stock: 10},
	{ID: 2, Name: "Oreo", Harga: 3000, Stock:30},
	{ID: 3, Name: "Wafer", Harga: 12000, Stock:20},
}


func getProductByID(w http.ResponseWriter, r *http.Request){
idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

		for _, p := range product {
			if p.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(p)
				return
			}
		}

		http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProductByID(w http.ResponseWriter, r *http.Request){
	//get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
			http.Error(w, "Invalid Product ID", http.StatusBadRequest)
			return
		}

	// get data dari request
	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop product cari id ganti sesuai data dari request
	for i := range product {
		if product[i].ID == id {
			product[i].ID = updateProduct.ID
			product[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)	
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProductByID(w http.ResponseWriter, r *http.Request){
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// loop product cari id
	for i, p := range product{
		if p.ID == id {
			// bikin slice baru dengan data sebeluym dan sesudah index
			product = append(product[:i], product[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success delete",
			})	
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}



func main() {
	// GET /api/product/{id}
	// PUT /api/product/{id}
	// DELETE /api/product/{id}
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductByID(w, r)
		case "PUT":
			updateProductByID(w, r)
		case "DELETE":
			deleteProductByID(w,r)
		}
	})


	// GET /api/product
	//POST /api/product
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		case "POST":
			// baca data dari request
			var newProduct  Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil{
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}
			
			//masukin data ke dalam variable product
			newProduct.ID = len(product) + 1
			product = append(product, newProduct)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		}
	})



	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
			"message": "API Running",
		})
	})
	fmt.Println("running server port localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed running server")
	}
}