package main

import (
	"fmt"
	"foodMenu/Handlers"
	"net/http"
)

func main() {
	fmt.Println("server started")

	http.HandleFunc("/", Handlers.FoodHandler)

	http.ListenAndServe(":8080", nil)
}
