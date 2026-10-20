package database

var products []Product

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

	products = append(products, prd1)
	products = append(products, prd2)
}

func StoreProducts(product Product) Product {
	product.Id = len(products) + 1
	products = append(products, product)

	return product
}

func ListProduct() []Product {
	return products
}

func GetSingleProduct(pid int) *Product {
	for _, product := range products {
		if product.Id == pid {
			return &product
		}
	}

	return nil
}

func UpdateProduct(product Product) {
	for idx, prod := range products {
		if prod.Id == product.Id {
			products[idx] = product
		}
	}
}

func DeleteByID(pid int) {
	var firstList []Product

	for _, product := range products {
		if pid != product.Id {
			firstList = append(firstList, product)
		}
	}

	products = firstList
}
