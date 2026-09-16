package model

type CardType struct {
	SubTypes   []string `json:"subTypes"`
	SuperTypes []string `json:"superTypes"`
}
