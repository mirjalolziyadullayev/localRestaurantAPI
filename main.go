package main

import (
	"fmt"
	"foodMenu/Handlers"
	"net/http"
)

func main() {
	fmt.Println("server started")

	http.HandleFunc("/national", Handlers.FoodHandler)
	http.HandleFunc("/others", Handlers.OthersHandler)

	http.ListenAndServe(":8080", nil)
}
