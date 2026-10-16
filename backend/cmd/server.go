package cmd

import (
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/global_router"
	"github.com/altr-r/implementing-go-backend/handlers"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	globalMux := global_router.GlobalRouter(mux)
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", globalMux)

	if err != nil {
		fmt.Println("There was an error starting the server")
	}

}
