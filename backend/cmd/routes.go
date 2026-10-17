package cmd

import (
	"net/http"

	"github.com/altr-r/implementing-go-backend/handlers"
	"github.com/altr-r/implementing-go-backend/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middleware.Logger))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.Logger))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID), middleware.Logger))
}
