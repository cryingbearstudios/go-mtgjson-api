package model

import "cryingbear.net/go-mtgjson-api/util"

type BoosterPack struct {
	Contents map[string]int `json:"contents"`
	Weight   int            `json:"weight"`
}

type BoosterSheet struct {
	AllowDuplicates *bool          `json:"allowDuplicates,omitempty"`
	BalanceColors   *bool          `json:"balanceColors,omitempty"`
	Cards           map[string]int `json:"cards"`
	Fixed           *bool          `json:"fixed,omitempty"`
	Foil            bool           `json:"foil"`
	TotalWeight     int            `json:"totalWeight"`
}

type BoosterConfig struct {
	Boosters            []BoosterPack           `json:"boosters"`
	BoostersTotalWeight int                     `json:"boostersTotalWeight"`
	Name                *string                 `json:"name,omitempty"`
	Sheets              map[string]BoosterSheet `json:"sheets"`
	SourceSetCodes      []util.SetCode          `json:"sourceSetCodes"`
}
