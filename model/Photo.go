package model

import "personspace/model"

type Photo struct {
	ID        int
	PhotoUser model.User
	Address   string
}
