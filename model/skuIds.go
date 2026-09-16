package model

type SkuIds struct {
	Etched  *string `json:"etched,omitempty"`
	Foil    *string `json:"foil,omitempty"`
	Nonfoil *string `json:"nonfoil,omitempty"`
	Other   *string `json:"other,omitempty"`
	Signed  *string `json:"signed,omitempty"`
}
