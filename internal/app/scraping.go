package app

import "github.com/brendsanchez/ws-money-go/internal/dto"

type Dollar interface {
	GetPrices() (*[]dto.Dollar, error)
}
