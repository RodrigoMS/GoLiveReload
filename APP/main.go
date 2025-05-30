package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

var mu sync.Mutex
var ultimaReinicio time.Time

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Servidor em execução !!!")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    fmt.Fprintf(w, "%v", ultimaReinicio.Unix())
}

func main() {
    mu.Lock()
    ultimaReinicio = time.Now()
    mu.Unlock()

    http.HandleFunc("/", handler)
    http.HandleFunc("/status", statusHandler)

    fmt.Println("Servidor rodando na porta 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
