package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/altr-r/implementing-go-backend/config"
	"github.com/altr-r/implementing-go-backend/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("There was an error starting the server")
		os.Exit(1)
	}
}
