package model

import "github.com/cryingbearstudios/go-mtgjson-api/util"

type DeckSet struct {
	Code               util.SetCode  `json:"code"`
	Commander          []CardSetDeck `json:"commander,omitempty"`
	MainBoard          []CardSetDeck `json:"mainBoard"`
	Name               string        `json:"name"`
	ReleaseDate        string        `json:"releaseDate"`
	SealedProductUuids []string      `json:"sealedProductUuids,omitempty"`
	SideBoard          []CardSetDeck `json:"sideBoard"`
	Type               string        `json:"type"`
}
