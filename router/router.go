package router

import (
	"github.com/gorilla/mux"

	"DroneNavigationSystem/location"
)

// InitializeRouter - mux router initialization
func InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	location.InitializeLocationAPI(r)
	return r
}

