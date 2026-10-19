package cmd

import (
	"github.com/altr-r/implementing-go-backend/config"
	"github.com/altr-r/implementing-go-backend/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)
}
