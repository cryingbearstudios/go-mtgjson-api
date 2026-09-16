package model

type SourceProducts struct {
	Etched  *[]string `json:"etched,omitempty"`
	Foil    *[]string `json:"foil,omitempty"`
	Nonfoil *[]string `json:"nonfoil,omitempty"`
}
