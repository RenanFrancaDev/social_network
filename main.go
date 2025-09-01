package main

import (
	"api/src/config"
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.HandleConfig()

	r := routes.HandleRoutes()
	fmt.Printf("Server running on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
