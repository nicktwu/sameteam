package models

const Male, Female, Nonbin = 0, 1, 2
const None, Dog, Cat, Other_sm, Other_lg, Many = 0, 1, 2, 3, 4, 5

type User struct {
	// Account Info
	Username     string `json:"username" bson:"username"` // eg: nicktwu
	Password     string `json:"password" bson:"-"`        // eg: ********
	PasswordHash []byte `json:"hash" bson:"hash"`         // eg: 1b82yhr!o&AS1dia
	Available    bool   `json:"avail" bson:"avail"`       // eg: true
	// Basic Biographical Info
	Name              string `json:"name" bson:"name"`           // eg: "Nick WooOOOooo'
	Phone             int    `json:"phone" bson:"phone"`         // eg: 1234567890
	Email             string `json:"email" bson:"email"`         // eg: "nwu@email.com"
	Institution       string `json:"institute" bson:"institute"` // eg: "MIT"
	Sex               int    `json:"sex" bson:"sex"`             // eg: 0 (male)
	Pets              int    `json:"pets" bson:"pets"`           // eg: 0 (none)
	Cleanliness       int    `json:"clean" bson:"clean"`         // eg: 4 (1-5)
	Wakeup            int    `json:"wakeup" bson:"wakeup"`       // eg: 0800 (0000-2400)
	Bedtime           int    `json:"bedtime" bson:"bedtime"`     // eg: 0330 (0000-2400)
	Smoking           bool   `json:"smoking" bson:"smoking"`     // eg: false
	PersonalStatement string `json:"pstate" bson:"pstate"`       // eg: "I love all-nighters"
	// Social History & Info
	VerifiedRating float32 `json:"rating" bson:"rating"`         // eg: 5.0 (0.0-5.0)
	KnownUsers     []User  `json:"knownusers" bson:"knownusers"` // eg: [Justin, Luis, Jessie, Alex]
	PastRoommates  []User  `json:"history" bson:"history"`       // eg: [TC, Eric]
	// Saved Preferences
	PrefSex   int  `json:"prefsex" bson:"prefsex"`     // eg: 0 (male)
	PrefPets  int  `json:"prefpets" bson:"prefpets"`   // eg: 0 (none)
	PrefClean int  `json:"prefclean" bson:"prefclean"` // eg: 8 (1-10)
	PrefWake  int  `json:"prefwake" bson:"prefwake"`   // eg: 1021 (0000-2400)
	PrefBed   int  `json:"prefbed" bson:"prefbed"`     // eg: 0520 (0000-2400)
	PrefSmoke bool `json:"prefsmoke" bson:"prefsmoke"` // eg: false
}
