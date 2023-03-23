package controllers

type Users struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transactions struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}
