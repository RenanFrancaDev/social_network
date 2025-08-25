package routes

import "net/http"

var userRoutes = []Router{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequireAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequireAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequireAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequireAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequireAuth: false,
	},
}
