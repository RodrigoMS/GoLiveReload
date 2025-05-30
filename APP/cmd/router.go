package server

import (
	"net/http"

	"github.com/RodrigoMS/GoLiveReload/cmd/controllers"
)

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controllers.HomePage)

	router.HandleFunc("GET /status", controllers.HomePage)

	return router
}