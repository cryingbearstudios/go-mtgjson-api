package model

type PurchaseUrls struct {
	CardKingdom              *string `json:"cardKingdom,omitempty"`
	CardKingdomEtched        *string `json:"cardKingdomEtched,omitempty"`
	CardKingdomFoil          *string `json:"cardKingdomFoil,omitempty"`
	Cardmarket               *string `json:"cardmarket,omitempty"`
	CardmarketFoil           *string `json:"cardmarketFoil,omitempty"`
	Tcgplayer                *string `json:"tcgplayer,omitempty"`
	TcgplayerAlternativeFoil *string `json:"tcgplayerAlternativeFoil,omitempty"`
	TcgplayerEtched          *string `json:"tcgplayerEtched,omitempty"`
}
