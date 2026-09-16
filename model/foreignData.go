package model

import "github.com/google/uuid"

type ForeignData struct {
	FaceName    *string     `json:"faceName,omitempty"`
	FlavorText  *string     `json:"flavorText,omitempty"`
	Identifiers Identifiers `json:"identifiers"`
	Language    string      `json:"language"`
	Name        string      `json:"name"`
	SkuIds      SkuIds      `json:"skuIds"`
	Text        *string     `json:"text,omitempty"`
	Type        *string     `json:"type,omitempty"`
	Uuid        uuid.UUID   `json:"uuid"`
}
