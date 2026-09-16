package model

import "github.com/google/uuid"

type Identifiers struct {
	AbuId                             *string    `json:"abuId,omitempty"`
	CardKingdomEtchedId               *string    `json:"cardKingdomEtchedId,omitempty"`
	CardKingdomFoilId                 *string    `json:"cardKingdomFoilId,omitempty"`
	CardKingdomId                     *string    `json:"cardKingdomId,omitempty"`
	CardSphereFoilId                  *string    `json:"cardsphereFoilId,omitempty"`
	CardSphereId                      *string    `json:"cardsphereId,omitempty"`
	CardtraderId                      *string    `json:"cardtraderId,omitempty"`
	CsiId                             *string    `json:"csiId,omitempty"`
	McmId                             *string    `json:"mcmId,omitempty"`
	McmMetaId                         *string    `json:"mcmMetaId,omitempty"`
	MiniatureMarketId                 *string    `json:"miniaturemarketId,omitempty"`
	MtgArenaId                        *string    `json:"mtgArenaId,omitempty"`
	MtgJsonFoilVersionId              *string    `json:"mtgjsonFoilVersionId,omitempty"`
	MtgJsonNonFoilVersionId           *string    `json:"mtgjsonNonFoilVersionId,omitempty"`
	MtgJsonV4Id                       *string    `json:"mtgjsonV4Id,omitempty"`
	MtgoFoilId                        *string    `json:"mtgoFoilId,omitempty"`
	MtgoId                            *string    `json:"mtgoId,omitempty"`
	MultiverseId                      *string    `json:"multiverseId,omitempty"`
	ScgId                             *string    `json:"scgId,omitempty"`
	ScryfallCardBackId                *string    `json:"scryfallCardBackId,omitempty"`
	ScryfallId                        *uuid.UUID `json:"scryfallId,omitempty"`
	ScryfallIllustrationId            *uuid.UUID `json:"scryfallIllustrationId,omitempty"`
	ScryfallOracleId                  *uuid.UUID `json:"scryfallOracleId,omitempty"`
	TcgplayerAlternativeFoilProductId *string    `json:"tcgplayerAlternativeFoilProductId,omitempty"`
	TcgplayerEtchedProductId          *string    `json:"tcgplayerEtchedProductId,omitempty"`
	TcgplayerProductId                *string    `json:"tcgplayerProductId,omitempty"`
	TntId                             *string    `json:"tntId,omitempty"`
}
