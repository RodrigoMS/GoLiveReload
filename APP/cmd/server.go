package server

import (
	"fmt"
	"net/http"

	"github.com/RodrigoMS/GoLiveReload/cmd/services"
)

func Start() {
	// Inicializa ultimaReinicio
  services.StartStatus()

	router := routes()

	fmt.Println("Servidor em execução na interface 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
			fmt.Println("Erro ao iniciar o servidor:", err)
	}
}