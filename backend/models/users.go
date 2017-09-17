package models

const Male, Female, Nonbin = 0, 1, 2
const None, Dog, Cat, Other_sm, Other_lg, Many = 0, 1, 2, 3, 4, 5

type User struct {
	Username          string `json:"username" bson:"username"`   // eg: nicktwu
	Password          []byte `json:"password" bson:"password"`   // eg: 1b82yhr!o&AS1dia
	Name              string `json:"name" bson:"name"`           // eg: "Nick WooOOOooo'
	Phone             string `json:"phone" bson:"phone"`         // eg: 123-456-7890
	Email             string `json:"email" bson:"email"`         // eg: "nwu@email.com"
	Sex               int    `json:"sex" bson:"sex"`             // eg: 0 (male)
	Pets              int    `json:"pets" bson:"pets"`           // eg: 0 (none)
	Cleanliness       int    `json:"clean" bson:"clean"`         // eg: 7 (1-10)
	Wakeup            int    `json:"wakeup" bson:"wakeup"`       // eg: 800 (0000-2400)
	Bedtime           int    `json:"bedtime" bson:"bedtime"`     // eg: 0330 (0000-2400)
	Smoking           bool   `json:"smoking" bson:"smoking"`     // eg: false
	PersonalStatement string `json:"pstate" bson:"pstate"`       // eg: "I love all-nighters"
	PrefSex           int    `json:"prefsex" bson:"prefsex"`     // eg: 0 (male)
	PrefPets          int    `json:"prefpets" bson:"prefpets"`   // eg: 0 (none)
	PrefClean         int    `json:"prefclean" bson:"prefclean"` // eg: 8 (1-10)
	PrefWake          int    `json:"prefwake" bson:"prefwake"`   // eg: 1021 (0000-2400)
	PrefBed           int    `json:"prefbed" bson:"prefbed"`     // eg: 0520 (0000-2400)
	PrefSmoke         bool   `json:"prefsmoke" bson:"prefsmoke"` // eg: false
}
