package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// SalvarCotacao salva a cotação em um arquivo no diretório `data/`
func SalvarArquivoCotacao(content string) error {
	// Definir diretório e arquivo
	dirPath := "data"
	filePath := filepath.Join(dirPath, "cotacao.txt")

	// Criar diretório se não existir
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("erro ao criar diretório: %w", err)
	}

	// Escrever conteúdo no arquivo
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("erro ao salvar cotação no arquivo: %w", err)
	}

	return nil
}
