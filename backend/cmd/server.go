package cmd

import (
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/global_router"
	"github.com/altr-r/implementing-go-backend/handlers"
	"github.com/altr-r/implementing-go-backend/middleware"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Logger)

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middleware.Logger))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.Logger))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID), middleware.Logger))

	globalMux := global_router.GlobalRouter(mux)
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", globalMux)

	if err != nil {
		fmt.Println("There was an error starting the server")
	}

}
