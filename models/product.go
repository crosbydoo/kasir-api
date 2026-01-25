package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"harga"`
	Stock int    `json:"stock"`
}

// penulisan Products daripada products
// karena Products = public/exported
// sedangkan products = private/unexported
var Products = []Product{
	{ID: 1, Name: "Indomie", Harga: 5000, Stock: 10},
	{ID: 2, Name: "Oreo", Harga: 3000, Stock: 30},
	{ID: 3, Name: "Wafer", Harga: 12000, Stock: 20},
	{ID: 4, Name: "Ultra Milk", Harga: 10000, Stock: 15},
}
