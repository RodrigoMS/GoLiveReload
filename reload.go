//Executar o VS Code no Slax
//code --no-sandbox --user-data-dir ~/vscode-root

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Mapa para armazenar timestamps dos arquivos
var arquivosModificados = make(map[string]time.Time)
var processoAtual *exec.Cmd // Variável para armazenar o processo da aplicação

func verificarAlteracoes(caminho string) {
	filepath.Walk(caminho, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ultimaModificacao, existe := arquivosModificados[path]
			if !existe || info.ModTime().After(ultimaModificacao) {
				fmt.Println("Arquivo modificado:", path)
				arquivosModificados[path] = info.ModTime()
				recompilarEExecutar()
			}
		}
		return nil
	})
}

func recompilarEExecutar() {
	fmt.Println("Recompilando aplicação...")
	cmdBuild := exec.Command("go", "build", "-o", "../app")
	cmdBuild.Dir = "APP"
	output, err := cmdBuild.CombinedOutput()
	if err != nil {
		fmt.Printf("Erro ao recompilar: %v\nSaída:\n%s\n", err, output)
		return
	}
	fmt.Println("Aplicação recompilada com sucesso!")

	// Se houver um processo rodando, encerrá-lo antes de iniciar um novo
	if processoAtual != nil {
		fmt.Println("Encerrando aplicação anterior...")
		err := processoAtual.Process.Kill() // Mata o processo anterior
		if err != nil {
			fmt.Printf("Erro ao encerrar aplicação: %v\n", err)
		}
		processoAtual = nil
	}

	fmt.Println("Iniciando nova aplicação...")
	cmdRun := exec.Command("./app") // No Windows, use "app.exe"
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	err = cmdRun.Start()
	if err != nil {
		fmt.Printf("Erro ao iniciar aplicação: %v\n", err)
	} else {
		processoAtual = cmdRun
		fmt.Println("Aplicação iniciada com sucesso!")
	}
}


func main() {
	caminho := "APP/"
	fmt.Println("Observando alterações em:", caminho)

	for {
		verificarAlteracoes(caminho)
		time.Sleep(2 * time.Second) // Ajuste o tempo conforme necessário
	}
}
