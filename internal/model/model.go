package model

type Car struct {
	Id       string  `json:"id"`
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Package  string  `json:"package"`
	Color    string  `json:"color"`
	Year     int     `json:"year"`
	Category string  `json:"category"`
	Mileage  float64 `json:"mileage"`
	Price    float64 `json:"price"`
}
