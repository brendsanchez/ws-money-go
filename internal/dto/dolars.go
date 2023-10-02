package dto

import "time"

type Dollar struct {
	Name      string
	Buy       *Price     `json:"buy,omitempty"`
	Sell      *Price     `json:"sell,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type ByName []Dollar

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

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
