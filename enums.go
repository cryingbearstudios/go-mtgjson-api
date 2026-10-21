package mtgJson

import "github.com/cryingbearstudios/go-mtgjson-api/util"

const (
	PrintingUnknown util.Printing = ""
	Normal          util.Printing = "NON FOIL"
	Foil            util.Printing = "FOIL"
)

const (
	LanguageUnknown    util.Language = ""
	English            util.Language = "ENGLISH"
	ChineseSimplified  util.Language = "CHINESE SIMPLIFIED"
	ChineseTraditional util.Language = "CHINESE TRADITIONAL"
	German             util.Language = "GERMAN"
	Italian            util.Language = "ITALIAN"
	Japanese           util.Language = "JAPANESE"
	Korean             util.Language = "KOREAN"
	Russian            util.Language = "RUSSIAN"
	French             util.Language = "FRENCH"
	PortugueseBrazil   util.Language = "PORTUGUESE BRAZIL"
	Spanish            util.Language = "SPANISH"
)

const (
	FinishUnknown util.Finish = ""
	Etched        util.Finish = "ETCHED"
)

const (
	ConditionUnknown util.Condition = ""
	NearMint         util.Condition = "NEAR MINT"
	LightlyPlayed    util.Condition = "LIGHTLY PLAYED"
	ModeratelyPlayed util.Condition = "MODERATELY PLAYED"
	HeavilyPlayed    util.Condition = "HEAVILY PLAYED"
	Damaged          util.Condition = "DAMAGED"
	Unopened         util.Condition = "UNOPENED"
)
