package entity

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price int    `json:"price"`
}
