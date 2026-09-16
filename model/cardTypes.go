package model

type CardTypes struct {
	Artifact     CardType `json:"artifact"`
	Battle       CardType `json:"battle"`
	Conspiracy   CardType `json:"conspiracy"`
	Creature     CardType `json:"creature"`
	Enchantment  CardType `json:"enchantment"`
	Instant      CardType `json:"instant"`
	Land         CardType `json:"land"`
	Phenomenon   CardType `json:"phenomenon"`
	Plane        CardType `json:"plane"`
	Planeswalker CardType `json:"planeswalker"`
	Scheme       CardType `json:"scheme"`
	Sorcery      CardType `json:"sorcery"`
	Tribal       CardType `json:"tribal"`
	Vanguard     CardType `json:"vanguard"`
}
