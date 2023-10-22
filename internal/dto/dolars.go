package dto

import "time"

type Dollar struct {
	Name      string     `json:"name,omitempty"`
	Buy       *Price     `json:"buy"`
	Sell      *Price     `json:"sell"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type Page struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
	Image string `json:"image"`
}

type ByName []Dollar

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type Price struct {
	ValueText string   `json:"valueText,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}

type DollarResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Health struct {
	Status string `json:"status" example:"ok"`
}
