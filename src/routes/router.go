package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Configure routes into router (mux)
func Configure(r *mux.Router) *mux.Router {
	routes := append(userRoutes, accountRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
