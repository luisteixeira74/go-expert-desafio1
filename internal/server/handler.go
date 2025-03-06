package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/luisteixeira74/go-expert-desafio1/internal/models"
)

// Handler para processar requisições do cliente
func HandleCotacao(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Criar um contexto com timeout de 200ms para a requisição HTTP
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		apiURL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
		// Criar a requisição HTTP com o contexto
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
		if err != nil {
			http.Error(w, "Erro ao criar requisição", http.StatusInternalServerError)
			return
		}

		// Executar a requisição HTTP
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao obter cotação externa", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Erro ao obter cotação externa", http.StatusInternalServerError)
			log.Println("Erro ao obter cotação externa")
			return
		}

		// Decodificar a resposta JSON
		var response map[string]models.Cotacao
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
			return
		}

		// Verificar se a cotação foi encontrada
		cotacao, exists := response["USDBRL"]
		if !exists {
			http.Error(w, "Cotação não encontrada", http.StatusNotFound)
			return
		}

		// Criar um contexto com timeout de 10ms para a inserção no banco
		ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelDB()

		// Inserir a cotação no banco de dados
		if err := cotacao.InserirCotacao(ctxDB, db); err != nil {
			http.Error(w, "Erro ao inserir cotação no banco", http.StatusInternalServerError)
			return
		}

		// Respondendo ao cliente
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	}
}
