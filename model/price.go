package model

type PriceSource string
type MtgoPriceSource PriceSource
type PaperPriceSource PriceSource

const (
	CardHoarder MtgoPriceSource  = "cardhoarder"
	CardKingdom PaperPriceSource = "cardkingdom"
	CardMarket  PaperPriceSource = "cardmarket"
	CardSphere  PaperPriceSource = "cardsphere"
	TCGPlayer   PaperPriceSource = "tcgplayer"
)

type PriceFormats struct {
	Mtgo  map[MtgoPriceSource]PriceList  `json:"mtgo,omitempty"`
	Paper map[PaperPriceSource]PriceList `json:"paper,omitempty"`
}

type PriceList struct {
	Buylist  *PricePoints `json:"buylist,omitempty"`
	Currency string       `json:"currency"`
	Retail   *PricePoints `json:"retail,omitempty"`
}

type PricePoints struct {
	Etched map[string]float64 `json:"etched,omitempty"`
	Foil   map[string]float64 `json:"foil,omitempty"`
	Normal map[string]float64 `json:"normal,omitempty"`
}
