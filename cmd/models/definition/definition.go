package definition

// WorldDef :
type WorldDef struct {
	Code             string     `firestore:"code"`
	Label            string     `firestore:"label"`
	Icon             string     `firestore:"icon"`
	Priority         int        `firestore:"priority"`
	Duration         int        `firestore:"duration"`
	DaysBeforeExpiry int        `firestore:"daysBeforeExpiry"`
	ShowExpiryDate   bool       `firestore:"showExpiryDate"`
	Levels           []LevelDef `firestore:"levels"`
}

// LevelDef :
type LevelDef struct {
	Code       string         `firestore:"code"`
	Label      string         `firestore:"label"`
	Order      int            `firestore:"order"`
	Prize      PrizeDef		  `firestore:"prize"`
	Challenges []ChallengeDef `firestore:"challenges"`
}

type PrizeDef struct {
	Category 	string `firestore:"category"`
	Label 		string `firestore:"label"`
	Value       int    `firestore:"value"`
}

// ChallengeDef :
type ChallengeDef struct {
	Code    string      `firestore:"code"`
	Label   string      `firestore:"label"`
	Actions []ActionDef `firestore:"actions"`
}

// ActionDef :
type ActionDef struct {
	EventId              string `firestore:"eventId"`
	Origin               string `firestore:"origin"`
	EventOnDifferentDays bool   `firestore:"eventOnDifferentDays"`
	Amount               int    `firestore:"amount"`
	Count                int    `firestore:"count"`
	CalculationType      string `firestore:"calculationType"`
}