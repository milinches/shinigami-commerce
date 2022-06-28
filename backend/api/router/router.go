package router

import (
	"github.com/gorilla/mux"
	"github.com/milinches/edsu-token-backend/api/router/routes"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}