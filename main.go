package main

import (
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting API")

	r := routes.HandleRoutes()

	log.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
