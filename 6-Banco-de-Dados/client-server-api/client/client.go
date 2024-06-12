package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const serverURL = "http://localhost:8080/cotacao"
const timeout = 300 * time.Millisecond

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Realiza a solicitação HTTP para obter a cotação do dólar do servidor
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Decodifica a resposta JSON
	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Fatal("Error decoding response:", err)
	}

	// Salva a cotação em um arquivo "cotacao.txt"
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "Dólar: %s\n", cotacao.Bid); err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Cotação do dólar salva em cotacao.txt com sucesso.")
}
