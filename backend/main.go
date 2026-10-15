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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)

	mux.HandleFunc("/get-product", getProducts)

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
