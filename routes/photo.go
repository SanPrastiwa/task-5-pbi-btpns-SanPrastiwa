package routes

import (
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/controllers/photocontroller"
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/middleware"
	"github.com/gorilla/mux"
)

func PhotoRoutes(router *mux.Router) {
	photo := router.PathPrefix("/photos").Subrouter()

	// Middleware
	photo.Use(middleware.Auth)

	photo.HandleFunc("", photocontroller.ListPhoto).Methods("GET")
	photo.HandleFunc("", photocontroller.CreatePhoto).Methods("POST")
	photo.HandleFunc("/{photoId}", photocontroller.ShowDetailPhoto).Methods("GET")
	photo.HandleFunc("/{photoId}", photocontroller.UpdatePhoto).Methods("PUT")
	photo.HandleFunc("/{photoId}", photocontroller.DeletePhoto).Methods("DELETE")
}
