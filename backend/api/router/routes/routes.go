package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milinches/edsu-token-backend/api/controllers"
)

var routes = []Route {
	{
		URI: "/products",
		Method: "GET",
		Handler: controllers.GetProducts,
	},
	{
		URI: "/product/{id}",
		Method: "GET",
		Handler: controllers.GetProduct,
	},
}

type (
	Route struct {
		URI     string
		Method  string
		Handler func(http.ResponseWriter, *http.Request)
	}
)

func Load() []Route {
	route := routes
	return route
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}