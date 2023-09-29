package dto

type Dollar struct {
	Name string
	Buy  Price
	Sell Price
}

type Price struct {
	Label string `json:"label"`
	Val   string `json:"value"`
}

type DollarResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"status,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Health struct {
	Status string `json:"status" example:"ok"`
}
