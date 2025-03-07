package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/luisteixeira74/go-expert-desafio1/internal/utils"
)

type valorDolar struct {
	Valor string `json:"bid"`
}

func main() {
	// Criando um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Criando a requisição HTTP com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	//O json é convertido em variavel array de bytes
	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}

	// Decodifica o JSON
	valorDolar := valorDolar{string(byteValue)}

	fmt.Println("Valor do Dolar:", valorDolar.Valor)

	// Monta o conteúdo do arquivo
	content := fmt.Sprintf("Dólar: %s\n", valorDolar.Valor)

	// Chama a função para salvar no arquivo
	if err := utils.SalvarArquivoCotacao(content); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cotação salva com sucesso em cotacao.txt")
}
