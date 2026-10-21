package model

import (
	"github.com/cryingbearstudios/go-mtgjson-api/util"
)

type TcgplayerSKUs struct {
	Condition util.Condition `json:"condition"`
	Finish    util.Finish    `json:"finish"`
	Language  util.Language  `json:"language"`
	Printing  util.Printing  `json:"printing"`
	ProductId int            `json:"productId"`
	SkuId     int            `json:"skuId"`
}
