package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisteixeira74/go-expert-desafio1/internal/database"
	"github.com/luisteixeira74/go-expert-desafio1/internal/server"
)

func main() {
	db, _ := database.ConnectDB()
	defer db.Close()

	http.HandleFunc("/cotacao", server.HandleCotacao(db))
	port := "8080"
	fmt.Println("Servidor rodando em http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
