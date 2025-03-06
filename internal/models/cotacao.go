package models

import (
	"context"
	"database/sql"
	"log"
)

// Estrutura para armazenar a cotação do dólar
type Cotacao struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	Bid        string `json:"bid"`
	CreateDate string `json:"create_date"`
}

// Método para inserir uma cotação no banco
func (c *Cotacao) InserirCotacao(ctx context.Context, db *sql.DB) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes (code, codein, name, bid, create_date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("Erro ao preparar statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, c.Code, c.CodeIn, c.Name, c.Bid, c.CreateDate)
	if err != nil {
		log.Printf("Erro ao inserir cotação: %v", err)
		return err
	}

	log.Println("Cotação inserida com sucesso!")
	return nil
}
