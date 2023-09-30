package dto

import "time"

type Dollar struct {
	Name      string
	Buy       *Price    `json:"buy,omitempty"`
	Sell      *Price    `json:"sell,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type Price struct {
	ValText string   `json:"valText,omitempty"`
	Val     *float64 `json:"val,omitempty"`
}

type DollarResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Health struct {
	Status string `json:"status" example:"ok"`
}
