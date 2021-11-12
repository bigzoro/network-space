package model

import "personspace/model"

type Document struct {
	ID           int
	DocumentUser model.User
	Address      string
}
