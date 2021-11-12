package model

import "personspace/model"

// 精彩时刻
type Moment struct {
	ID              int
	MomentUser      model.User
	Text            string
	PhotoAddress    string
	DocumentAddress string
}
