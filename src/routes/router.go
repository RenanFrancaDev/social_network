package routes

import (
	"api/src/middlewares"
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

		if route.RequireAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authentication(route.Function))).Methods(route.Method)
		}

		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return r
}
