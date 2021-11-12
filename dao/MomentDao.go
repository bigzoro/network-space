package dao

import (
	"personSpace/model"
)

type MomentDao struct {
}

// 添加精彩时刻
func (dd *DiaryDao) addMoment(diary model.Diary) bool {

	return true
}

// 删除精彩时刻
func (dd *DiaryDao) delteteMoment(id int) bool {

	return true
}

// 修改精彩时刻
func (dd *DiaryDao) updateMoment(id int, diary model.Diary) bool {

	return true
}

// 查询精彩时刻
func (dd *DiaryDao) queryMoment(id int) bool {

	return true
}
