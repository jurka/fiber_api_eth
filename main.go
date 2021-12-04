package main

import (
	handlers "fiber_api_eth/handlers"
	_ "fiber_api_eth/models"
	"fiber_api_eth/routes"
	"log"
	// _ "swagger/docs"
)

func main() {
	app, err := handlers.NewApp("https://ropsten.infura.io/v3/ebaab4889720457280559f4f622756df", "0x4f7f1380239450AAD5af611DB3c3c1bb51049c29")
	if err != nil {
		log.Fatalln(err)
	}

	router := routes.New(app)
	router.Listen(":9090")

}
