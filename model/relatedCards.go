package model

type RelatedCards struct {
	ReverseRelated *[]string `json:"reverseRelated,omitempty"`
	Spellbook      *[]string `json:"spellbook,omitempty"`
	Tokens         *[]string `json:"tokens,omitempty"`
}
