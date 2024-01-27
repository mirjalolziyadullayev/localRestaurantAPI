package models 

type FoodModel struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Included_Drinks []string `json:"included_drinks"`
}