package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/config"
	"github.com/SanPrastiwa/task-5-pbi-btpns-SanPrastiwa/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	router := mux.NewRouter()
	routes.RoutesIndex(router)

	log.Println("[APP] Server is listening to port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), router)
}
