package routes

import (
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/controllers/usercontroller"
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/middleware"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	user := router.PathPrefix("/users").Subrouter()

	// Middleware
	user.Use(middleware.Auth)

	user.HandleFunc("", usercontroller.ListUser).Methods("GET")
	user.HandleFunc("/{userId}", usercontroller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{userId}", usercontroller.DeleteUser).Methods("DELETE")
}
