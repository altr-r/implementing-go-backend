package database

var Products []Product

type Product struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Descripttion string  `json:"description"`
	Price        float64 `json:"price"`
	ImgUrl       string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		Id:           1,
		Title:        "Apple",
		Descripttion: "This is an apple",
		Price:        10,
		ImgUrl:       "Wow",
	}

	prd2 := Product{
		Id:           2,
		Title:        "Orange",
		Descripttion: "This is an orange",
		Price:        12,
		ImgUrl:       "WOw2",
	}

	Products = append(Products, prd1)
	Products = append(Products, prd2)
}
