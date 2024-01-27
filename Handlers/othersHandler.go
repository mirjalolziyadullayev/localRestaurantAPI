package Handlers

import (
	"encoding/json"
	"foodMenu/models"
	"net/http"
	"os"
)

func OthersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getOtherFoods(w, r)
	}
}
func getOtherFoods(w http.ResponseWriter, r *http.Request) {

	var Data []models.FoodModel
	ByteCollection, _ := os.ReadFile("db/others.json")
	json.Unmarshal(ByteCollection, &Data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Data)
}
