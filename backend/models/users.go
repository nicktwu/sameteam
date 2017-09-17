package models

const Male, Female, Nonbin = 0, 1, 2
const None, Dog, Cat, Other_sm, Other_lg, Many = 0, 1, 2, 3, 4, 5

type User struct {
	Name              string // eg: "Nick WooOOOooo'
	Phone             string // eg: 123-456-7890
	Username          string // eg: nicktwu
	Password          []byte // eg: 1b82yhr!o&AS1dia
	Sex               int    // eg: 0 (male)
	Pets              int    // eg: 0 (none)
	Cleanliness       int    // eg: 7 (1-10)
	Wakeup            int    // eg: 800 (0000-2400)
	Bedtime           int    // eg: 0330 (0000-2400)
	Smoking           bool   // eg: false
	PersonalStatement string // eg: "I love all-nighters"
	PrefSex           int    // eg: 0 (male)
	PrefPets          int    // eg: 0 (none)
	PrefClean         int    // eg: 8 (1-10)
	PrefWake          int    // eg: 1021 (0000-2400)
	PrefBed           int    // eg: 0520 (0000-2400)
	PrefSmoke         bool   // eg: false
}
