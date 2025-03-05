package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Estrutura para armazenar a cotação do dólar
type Cotacao struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	Bid        string `json:"bid"`
	CreateDate string `json:"create_date"`
}

// Handler para processar requisições do cliente
func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		http.Error(w, "Erro ao chamar API externa", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler resposta da API", http.StatusInternalServerError)
		return
	}

	var data map[string]Cotacao
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Erro ao converter JSON", http.StatusInternalServerError)
		return
	}

	// Pegando o campo necessário (USD)
	cotacao, exists := data["USDBRL"]
	if !exists {
		http.Error(w, "Dados não encontrados", http.StatusNotFound)
		return
	}

	// Respondendo ao cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cotacao)
}
