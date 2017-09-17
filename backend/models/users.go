package models

const Male, Female, Nonbin = 0, 1, 2
const Dog, Cat, Other_sm, Other_lg = 0, 1, 2, 3

type User struct {
	Name              string
	Phone             string
	Sex               int
	Pets              int
	Cleanliness       int
	Wakeup            int
	Bedtime           int
	Smoking           bool
	PersonalStatement string
}
