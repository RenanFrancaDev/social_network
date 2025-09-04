package routes

import (
	"github.com/gorilla/mux"
)

func HandlerRoutes() *mux.Router {
	r := mux.NewRouter()
	return Configure(r)
}
