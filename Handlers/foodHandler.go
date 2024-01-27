package Handlers

import (
	"encoding/json"
	"foodMenu/models"
	"net/http"
	"os"
)

func FoodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getFoods(w, r)
	}
}
func getFoods(w http.ResponseWriter, r *http.Request) {

	var Data []models.FoodModel
	ByteCollection, _ := os.ReadFile("db/all.json")
	json.Unmarshal(ByteCollection, &Data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Data)
}
