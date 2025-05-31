package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	fileModTimes = make(map[string]time.Time)
	currentProc  *exec.Cmd
	appPath      = "APP/"
	buildOutput  = "../app"
)

func checkChanges() bool {
	changed := false

	filepath.Walk(appPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if lastMod, exists := fileModTimes[path]; !exists || info.ModTime().After(lastMod) {
			fmt.Printf("Arquivo modificado: %s\n", path)
			fileModTimes[path] = info.ModTime()
			changed = true
		}
		return nil
	})

	return changed
}

func rebuildAndRun() {
	fmt.Println("Recompilando aplicação...")
	cmd := exec.Command("go", "build", "-o", buildOutput)
	cmd.Dir = appPath
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Erro ao recompilar: %v\nSaída:\n%s\n", err, output)
		return
	}

	if currentProc != nil {
		fmt.Println("Encerrando aplicação anterior...")
		_ = currentProc.Process.Kill()
	}

	fmt.Println("Iniciando nova aplicação...")
	currentProc = exec.Command("./app")
	currentProc.Stdout = os.Stdout
	currentProc.Stderr = os.Stderr
	
	if err := currentProc.Start(); err != nil {
		fmt.Printf("Erro ao iniciar aplicação: %v\n", err)
		currentProc = nil
	}
}

func main() {
	fmt.Printf("Observando alterações em: %s\n", appPath)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if checkChanges() {
			rebuildAndRun()
		}
	}
}