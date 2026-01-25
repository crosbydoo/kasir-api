package models

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Category Makanan"},
	{ID: 2, Name: "Minuman", Description: "Category Minuman"},
}
