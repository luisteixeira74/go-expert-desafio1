package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luisteixeira74/go-expert-desafio1/internal/server"
)

func main() {
	http.HandleFunc("/cotacao", server.HandleCotacao)

	port := "8080"
	fmt.Println("Servidor rodando em http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
