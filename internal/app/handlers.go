package app

import "net/http"

type Handlers interface {
	GetPrices() http.HandlerFunc
	GetPages() http.HandlerFunc
}
