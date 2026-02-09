package models

type Transaction struct{
	ID int `json:"id"`
	TotalAmount int `json:"total_amount"`
	Details []TransactionDetail `json:"details"` 
}

type TransactionDetail struct{
	ID int `json:"id"`
	ProductID int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	ProductName string `json:"product_name"`
	Quantity int `json:"quantity"`
	Subtotal int `json:"subtotal"`
}

type CheckoutRequest struct{
	Items []CheckoutItems `json:"items"`
}

type CheckoutItems struct{
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}