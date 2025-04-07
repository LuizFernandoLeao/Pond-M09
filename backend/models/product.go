package models

type Product struct {
    ID          int     `json:"id"`
    Image       string  `json:"image"`
    Price       float64 `json:"price"`
    Quantity    int     `json:"quantity"`
    Description string  `json:"description"`
}