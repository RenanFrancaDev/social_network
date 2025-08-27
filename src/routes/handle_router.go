package routes

import (
	"github.com/gorilla/mux"
)

func HandleRoutes() *mux.Router {
	r := mux.NewRouter()
	return Configure(r)
}
