package model

import "cryingbear.net/go-mtgjson-api/util"

type SetList struct {
	BaseSetSize      int             `json:"baseSetSize"`
	Block            *string         `json:"block,omitempty"`
	CardSphereSetId  *int            `json:"cardsphereSetId,omitempty"`
	Code             util.SetCode    `json:"code"`
	Decks            []DeckSet       `json:"decks,omitempty"`
	IsFoilOnly       bool            `json:"isFoilOnly"`
	IsForeignOnly    *bool           `json:"isForeignOnly,omitempty"`
	IsNonFoilOnly    *bool           `json:"isNonFoilOnly,omitempty"`
	IsOnlineOnly     bool            `json:"isOnlineOnly"`
	IsPaperOnly      *bool           `json:"isPaperOnly,omitempty"`
	IsPartialPreview *bool           `json:"isPartialPreview,omitempty"`
	KeyruneCode      string          `json:"keyruneCode"`
	Languages        []string        `json:"languages,omitempty"`
	McmId            *int            `json:"mcmId,omitempty"`
	McmIdExtras      *int            `json:"mcmIdExtras,omitempty"`
	McmName          *string         `json:"mcmName,omitempty"`
	MtgoCode         *string         `json:"mtgoCode,omitempty"`
	Name             string          `json:"name"`
	ParentCode       *util.SetCode   `json:"parentCode,omitempty"`
	ReleaseDate      string          `json:"releaseDate"`
	SealedProduct    []SealedProduct `json:"sealedProduct,omitempty"`
	TcgplayerGroupId *int            `json:"tcgplayerGroupId,omitempty"`
	TokenSetCode     *util.SetCode   `json:"tokenSetCode,omitempty"`
	TotalSetSize     int             `json:"totalSetSize"`
	Translations     Translations    `json:"translations"`
	Type             string          `json:"type"`
}
