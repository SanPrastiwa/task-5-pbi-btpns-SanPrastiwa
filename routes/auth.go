package routes

import (
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/controllers/authcontroller"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	auth := router.PathPrefix("/users").Subrouter()

	auth.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	auth.HandleFunc("/register", authcontroller.Register).Methods("POST")
	auth.HandleFunc("/login", authcontroller.Login).Methods("POST")
}
