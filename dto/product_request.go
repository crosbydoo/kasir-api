package dto

type ProductRequest struct {
	Name  string `json:"name"`
	Harga int    `json:"harga"`
	Stock int    `json:"stock"`
}
