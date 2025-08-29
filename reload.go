package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var fileModTimes = make(map[string]time.Time)
var currentProc *exec.Cmd

func runApp() {
	fmt.Println("\nExecutando ...")

	if currentProc != nil {
		err := currentProc.Process.Kill()

		if err != nil {
			fmt.Printf("Erro ao encerrar o processo: %v\n", err)
		}
		currentProc.Wait()
	}

	currentProc = exec.Command("./app.exe")

	currentProc.Stdout = os.Stdout
	currentProc.Stderr = os.Stderr

	err := currentProc.Start()
	if err != nil {
		fmt.Printf("Erro ao iniciar app: %v\n", err)
		currentProc = nil
	}

	fmt.Println("\nPronto!")
}

func rebuildApp() {
	fmt.Println("\n=== Recompilando APP ===")

	cmd := exec.Command("go", "build", "-o", "../app.exe")
	cmd.Dir = "./APP"

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Erro ao gerar o APP: %v\n", err)
		fmt.Printf("Saída do compilador:\n%s\n", output)
	}
}

func checkChanges() bool {
	changed := false

	err := filepath.Walk("APP", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		lastMod, exists := fileModTimes[path]
		if !exists || info.ModTime().After(lastMod) {
			fileModTimes[path] = info.ModTime()
			changed = true
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
	}

	return changed
}

func main() {
	fmt.Println("Go Live Reload - Iniciando monitoramento...")
	fmt.Println("Monitorando pasta: APP/")
	fmt.Println("Pressione Ctrl C para sair")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-interrupt
		fmt.Println("\nDesligando ....")
		if currentProc != nil {
			currentProc.Process.Kill()
		}
		os.Exit(0)
	}()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if checkChanges() {
			fmt.Println("\n---------------------------------")
			rebuildApp()
			runApp()
		}
	}
}
