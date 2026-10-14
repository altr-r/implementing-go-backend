package main

import (
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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)

	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("There was an error starting the server")
	}
}
