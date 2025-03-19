package model

type OrderRequest struct {
	CustomerName    string           `json:"customer_name,omitempty"`
	DeliveryAddress string           `json:"delivery_address,omitempty"`
	BottlesCount    int32            `json:"bottles_count,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	TestData        []AdditionalInfo `json:"test_data,omitempty"`
}

type AdditionalInfo struct {
	TextField    string   `json:"text_field,omitempty"`
	BinaryData   []byte   `json:"binary_data,omitempty"`
	NumericValue float64  `json:"numeric_value,omitempty"`
	Timestamp    int64    `json:"timestamp,omitempty"`
	StringArray  []string `json:"string_array,omitempty"`
	IntArray     []int32  `json:"int_array,omitempty"`
	ImageData    []byte   `json:"image_data,omitempty"`
	LongText     string   `json:"long_text,omitempty"`
}

type Order struct {
	ID              string `json:"id"`
	CustomerName    string `json:"customer_name"`
	DeliveryAddress string `json:"delivery_address"`
	BottlesCount    int32  `json:"bottles_count"`
	PhoneNumber     string `json:"phone_number"`
}
