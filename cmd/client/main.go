package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Estrutura para armazenar a cotação recebida do servidor
type Cotacao struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	Bid        string `json:"bid"`
	CreateDate string `json:"create_date"`
}

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

	var cotacao Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	fmt.Printf("Cotação do dólar: %s BRL\n", cotacao.Bid)
}
