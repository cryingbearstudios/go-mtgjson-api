package model

import "cryingbear.net/go-mtgjson-api/util"

type Set struct {
	BaseSetSize      int                      `json:"baseSetSize"`
	Block            *string                  `json:"block,omitempty"`
	Booster          map[string]BoosterConfig `json:"booster,omitempty"`
	Cards            []CardSet                `json:"cards"`
	CardSphereSetId  *int                     `json:"cardsphereSetId,omitempty"`
	Code             util.SetCode             `json:"code"`
	Decks            []DeckSet                `json:"decks,omitempty"`
	IsFoilOnly       bool                     `json:"isFoilOnly"`
	IsForeignOnly    *bool                    `json:"isForeignOnly,omitempty"`
	IsNonFoilOnly    *bool                    `json:"isNonFoilOnly,omitempty"`
	IsOnlineOnly     bool                     `json:"isOnlineOnly"`
	IsPaperOnly      *bool                    `json:"isPaperOnly,omitempty"`
	IsPartialPreview *bool                    `json:"isPartialPreview,omitempty"`
	KeyruneCode      string                   `json:"keyruneCode"`
	Languages        []string                 `json:"languages,omitempty"`
	MCMId            *int                     `json:"mcmId,omitempty"`
	MCMIdExtras      *int                     `json:"mcmIdExtras,omitempty"`
	MCMName          *string                  `json:"mcmName,omitempty"`
	MTGOCode         *string                  `json:"mtgoCode,omitempty"`
	Name             string                   `json:"name"`
	ParentCode       *string                  `json:"parentCode,omitempty"`
	ReleaseDate      string                   `json:"releaseDate"`
	SealedProduct    []SealedProduct          `json:"sealedProduct,omitempty"`
	TCGPlayerGroupId *int                     `json:"tcgplayerGroupId,omitempty"`
	TokenSetCode     *util.SetCode            `json:"tokenSetCode,omitempty"`
	Tokens           []CardToken              `json:"tokens"`
	TotalSetSize     int                      `json:"totalSetSize"`
	Translations     Translations             `json:"translations"`
	Type             string                   `json:"type"`
}
