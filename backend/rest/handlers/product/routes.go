package product

import (
	"net/http"

	middleware "github.com/altr-r/implementing-go-backend/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), middleware.AuthenticateJwt))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct)))

}
