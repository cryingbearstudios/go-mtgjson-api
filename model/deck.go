package model

import "cryingbear.net/go-mtgjson-api/util"

type Deck struct {
	Code               util.SetCode `json:"code"`
	Commander          []CardDeck   `json:"commander,omitempty"`
	MainBoard          []CardDeck   `json:"mainBoard"`
	Name               string       `json:"name"`
	ReleaseDate        string       `json:"releaseDate"`
	SealedProductUuids []string     `json:"sealedProductUuids,omitempty"`
	SideBoard          []CardDeck   `json:"sideBoard"`
	Tokens             []CardToken  `json:"tokens,omitempty"`
	Type               string       `json:"type"`
}
