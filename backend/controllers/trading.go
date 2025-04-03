package controllers

import (
	"backend/services"
	"encoding/json"
	"net/http"
)

func GetPairs(w http.ResponseWriter, r *http.Request) {
	pairs := services.GetTradingPairs()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pairs)
}
