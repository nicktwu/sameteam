package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Roommate struct {
	Name  string
	Phone string
	// etc... form information
}
