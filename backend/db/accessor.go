package db

import (
	"github.com/nicktwu/sameteam/backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DB_NAME = "sameteam"
const USER_COLLECTION = "users"

func CheckUserExists(s *mgo.Session, username string) (bool, error) {
	count, err := s.DB(DB_NAME).C(USER_COLLECTION).Find(bson.M{"username": username}).Count()
	return count > 0, err
}

func InsertUser(s *mgo.Session, user models.User) error {
	return s.DB(DB_NAME).C(USER_COLLECTION).Insert(user)
}
