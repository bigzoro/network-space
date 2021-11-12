package model

import "personspace/model"

// 日记
// hello, how are you?
type Diary struct {
	ID              int
	DiaryUser       model.User
	Text            string
	PhotoAddress    string
	DocumentAddress string
}
