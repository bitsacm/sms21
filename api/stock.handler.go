package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dush-t/sms21/db/models"
	"github.com/gorilla/mux"
)

// AddStock adds a stock to the database
func AddStock(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var stock models.Stock
		err := json.NewDecoder(r.Body).Decode(&stock)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, addErr := data.Stocks.Create(stock)
		if addErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error creating stock:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		payload := struct {
			ID    string  `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}{id, stock.Name, stock.Price}
		json.NewEncoder(w).Encode(payload)
	})
}

//GetStockData fetches the data of a stock by stock ID
func GetStockData(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		stock, readErr := data.Stocks.GetByID(id)
		if readErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error reading stock:", readErr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		payload := struct {
			ID    string  `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}{stock.ID, stock.Name, stock.Price}
		json.NewEncoder(w).Encode(payload)
	})
}
