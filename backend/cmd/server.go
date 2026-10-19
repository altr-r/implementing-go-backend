package cmd

import (
	"fmt"
	"net/http"

	"github.com/altr-r/implementing-go-backend/middleware"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.CorsWithPreflight)

	initRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("There was an error starting the server")
	}
}
