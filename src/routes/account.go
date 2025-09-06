package routes

import (
	"api/src/controllers"
	"net/http"
)

var accountRoutes = []Router{
	{
		URI:         "/signin",
		Method:      http.MethodPost,
		Function:    controllers.SignIn,
		RequireAuth: false,
	},
}
