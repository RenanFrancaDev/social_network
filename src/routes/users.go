package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Router{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: true,
	},
	{
		URI:         "/searchUsers",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/user/{userID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
