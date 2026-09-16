package model

import (
	"cryingbear.net/go-mtgjson-api/util"
	"github.com/google/uuid"
)

type SealedProduct struct {
	CardCount    *int                   `json:"cardCount,omitempty"`
	Category     *string                `json:"category,omitempty"`
	Contents     *SealedProductContents `json:"contents,omitempty"`
	Identifiers  Identifiers            `json:"identifiers"`
	Name         string                 `json:"name"`
	ProductSize  *int                   `json:"productSize,omitempty"`
	PurchaseUrls PurchaseUrls           `json:"purchaseUrls"`
	ReleaseDate  *string                `json:"releaseDate,omitempty"`
	Subtype      *string                `json:"subtype,omitempty"`
	Uuid         uuid.UUID              `json:"uuid"`
}

type SealedProductContents struct {
	Card     []SealedProductCard     `json:"card,omitempty"`
	Deck     []SealedProductDeck     `json:"deck,omitempty"`
	Other    []SealedProductOther    `json:"other,omitempty"`
	Pack     []SealedProductPack     `json:"pack,omitempty"`
	Sealed   []SealedProductSealed   `json:"sealed,omitempty"`
	Variable []SealedProductContents `json:"variable,omitempty"`
}

type SealedProductCard struct {
	Foil   *bool     `json:"foil,omitempty"`
	Name   string    `json:"name"`
	Number string    `json:"number"`
	Set    string    `json:"set"`
	Uuid   uuid.UUID `json:"uuid"`
}

type SealedProductDeck struct {
	Name string `json:"name"`
	Set  string `json:"set"`
}

type SealedProductOther struct {
	Name string `json:"name"`
}

type SealedProductPack struct {
	Code util.SetCode `json:"code"`
	Set  string       `json:"set"`
}

type SealedProductSealed struct {
	Count int       `json:"count"`
	Name  string    `json:"name"`
	Set   string    `json:"set"`
	Uuid  uuid.UUID `json:"uuid"`
}
