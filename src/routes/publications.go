package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Router{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		Function:    controllers.CreatePublicaton,
		RequireAuth: true,
	},
	{
		URI:         "/publications",
		Method:      http.MethodGet,
		Function:    controllers.GetPublicatons,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationID}",
		Method:      http.MethodGet,
		Function:    controllers.GetPublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationID}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationID}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePublication,
		RequireAuth: true,
	},
}
