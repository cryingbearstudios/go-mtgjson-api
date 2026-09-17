package model

import "github.com/cryingbearstudios/go-mtgjson-api/util"

type DeckList struct {
	Code        util.SetCode `json:"code"`
	FileName    string       `json:"fileName"`
	Name        string       `json:"name"`
	ReleaseDate string       `json:"releaseDate"`
	Type        string       `json:"type"`
}
