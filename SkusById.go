package mtgJson

import (
	"github.com/cryingbearstudios/go-mtgjson-api/model"
	"github.com/google/uuid"
)

type SkusById struct {
	Id   uuid.UUID             `json:"id"`
	SKUs []model.TcgplayerSKUs `json:"skus"`
}
