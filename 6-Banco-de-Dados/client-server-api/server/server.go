package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	apiURL        = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	serverAddress = ":8080"
	fetchTimeout  = 200 * time.Millisecond
	saveTimeout   = 10 * time.Millisecond
)

type Cotacao struct {
	ID        uint      `gorm:"primaryKey"`
	Bid       string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}

type ApiResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func fetchCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return "", err
	}

	return apiResponse.USDBRL.Bid, nil
}

func saveCotacao(ctx context.Context, db *gorm.DB, bid string) error {
	cotacao := Cotacao{
		Bid:       bid,
		Timestamp: time.Now(),
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		result := db.WithContext(ctx).Create(&cotacao)
		return result.Error
	}
}

func cotacaoHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), fetchTimeout)
		defer cancel()

		bid, err := fetchCotacao(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusRequestTimeout)
			log.Println("Error fetching cotacao:", err)
			return
		}

		ctxSave, cancelSave := context.WithTimeout(context.Background(), saveTimeout)
		defer cancelSave()

		err = saveCotacao(ctxSave, db, bid)
		if err != nil {
			log.Println("Error saving cotacao:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"bid": bid})
	}
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../cotacoes.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Cotacao{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatal("Failed to set up database:", err)
	}

	http.HandleFunc("/cotacao", cotacaoHandler(db))

	fmt.Println("Server is running at", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
