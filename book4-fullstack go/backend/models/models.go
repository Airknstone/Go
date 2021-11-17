package models

type Product struct {
	Image       string  `json:"img"`
	ImageAlt    string  `json:"imgalt"`
	Price       float64 `json:"promotion"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productname"`
	Description string  `json:"desc"`
}
