package controllers

import (
	"fmt"
	"net/http"

	"github.com/RodrigoMS/GoLiveReload/cmd/services"
)

func StatusService(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", services.GetLastRestart())
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor em execução!!!")
}
