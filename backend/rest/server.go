package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/altr-r/implementing-go-backend/config"
	"github.com/altr-r/implementing-go-backend/rest/handlers/product"
	"github.com/altr-r/implementing-go-backend/rest/handlers/user"
	middleware "github.com/altr-r/implementing-go-backend/rest/middlewares"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("There was an error starting the server")
		os.Exit(1)
	}
}
