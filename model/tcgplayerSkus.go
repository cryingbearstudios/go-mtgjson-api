package model

type Condition string

const (
	NearMint         Condition = "NEAR MINT"
	LightlyPlayed    Condition = "LIGHTLY PLAYED"
	ModeratelyPlayed Condition = "MODERATELY PLAYED"
	HeavilyPlayed    Condition = "HEAVILY PLAYED"
	Damaged          Condition = "DAMAGED"
)

type TcgplayerSKUs struct {
	Condition Condition `json:"condition"`
	Finish    string    `json:"finish"`
	Language  string    `json:"language"`
	Printing  string    `json:"printing"`
	ProductId int       `json:"productId"`
	SkuId     int       `json:"skuId"`
}
