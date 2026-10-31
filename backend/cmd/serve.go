package cmd

import (
	"github.com/altr-r/implementing-go-backend/config"
	"github.com/altr-r/implementing-go-backend/rest"
	"github.com/altr-r/implementing-go-backend/rest/handlers/product"
	"github.com/altr-r/implementing-go-backend/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	server := rest.NewServer(cnf, product.NewHandler(), user.NewHandler())

	server.Start()
}
