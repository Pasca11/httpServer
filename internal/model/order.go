package model

type Order struct {
	ID          string `json:"id"`
	OfficeID    string `json:"office_id"`
	WaterAmount int    `json:"water_amount"`
	Status      string `json:"status"`
}
