package app

import "github.com/brendsanchez/ws-money-go/internal/dto"

type UseCase interface {
	GetPrices(webPage string) dto.DollarResponse[*[]dto.Dollar]
	GetPages() dto.DollarResponse[*[]dto.Page]
}
