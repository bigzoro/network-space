package dao

import (
	"personSpace/model"
)

type DiaryDao struct {
}

// 添加日记
func (dd *DiaryDao) addDiary(diary model.Diary) bool {

	return true
}

// 删除日记
func (dd *DiaryDao) delteteDiary(id int) bool {

	return true
}

// 修改日记
func (dd *DiaryDao) updateDiary(id int, diary model.Diary) bool {

	return true
}

// 查询日记
func (dd *DiaryDao) queryDiary(id int) bool {

	return true
}
