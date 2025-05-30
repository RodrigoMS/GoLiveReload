package services

import (
    "sync"
    "time"
)

var mu sync.Mutex
var ultimaReinicio time.Time

// Inicializa a variável ultimaReinicio
func StartStatus() {
    mu.Lock()
    defer mu.Unlock()
    ultimaReinicio = time.Now()
}

// Retorna o timestamp da última reinicialização
func GetLastRestart() int64 {
    mu.Lock()
    defer mu.Unlock()
    return ultimaReinicio.Unix()
}
