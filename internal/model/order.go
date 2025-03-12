package model

type Order struct {
	ID              string `json:"id"`
	CustomerName    string `json:"customer_name"`
	DeliveryAddress string `json:"delivery_address"`
	BottlesCount    int32  `json:"bottles_count"`
	PhoneNumber     string `json:"phone_number"`
}
