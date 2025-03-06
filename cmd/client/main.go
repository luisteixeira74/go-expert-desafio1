package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/luisteixeira74/go-expert-desafio1/internal/models"
	"github.com/luisteixeira74/go-expert-desafio1/internal/utils"
)

func main() {
	resp, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		fmt.Println("Erro ao chamar o servidor:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	if !json.Valid(body) {
		log.Println("A resposta não é um JSON válido:", string(body))
		return
	}

	// Decodifica o JSON
	var cotacao models.Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	// Monta o conteúdo do arquivo
	content := fmt.Sprintf("Dólar: %s\n", cotacao.Bid)

	// Chama a função para salvar no arquivo
	if err := utils.SalvarArquivoCotacao(content); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cotação salva com sucesso!")
}
