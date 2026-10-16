package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World"))
}

type Product struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Descripttion string  `json:"description"`
	Price        float64 `json:"price"`
	ImgUrl       string  `json:"imageUrl"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Not Get Request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(products)

	if err != nil {
		fmt.Println("Not a valid json")
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Acess-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Not Get Request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Not a valid JSON")
		http.Error(w, "Give valid json", 400)
		return
	}

	newProduct.Id = len(products) + 1

	products = append(products, newProduct)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)

	mux.HandleFunc("/get-products", getProducts)

	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("There was an error starting the server")
	}
}

func init() {
	prd1 := Product{
		Id:           1,
		Title:        "Apple",
		Descripttion: "This is an apple",
		Price:        10,
		ImgUrl:       "Wow",
	}

	prd2 := Product{
		Id:           2,
		Title:        "Orange",
		Descripttion: "This is an orange",
		Price:        12,
		ImgUrl:       "WOw2",
	}

	products = append(products, prd1)
	products = append(products, prd2)
}
