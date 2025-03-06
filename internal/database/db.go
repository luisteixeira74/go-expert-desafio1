package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Função para conectar ao banco SQLite
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/cotacoes.db")
	if err != nil {
		return nil, err
	}

	// Criar tabela se não existir
	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		bid TEXT,
		create_date TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	return db, nil
}
