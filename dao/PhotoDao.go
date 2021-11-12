package dao

import (
	"personSpace/model"
)

type PhotoDao struct {
}

// 添加图片
func (dd *PhotoDao) addPhoto(Photo model.Photo) bool {

	return true
}

// 删除图片
func (dd *PhotoDao) deltetePhoto(id int) bool {

	return true
}

// 修改图片
func (dd *PhotoDao) updatePhoto(id int, Photo model.Photo) bool {

	return true
}

// 查询图片
func (dd *PhotoDao) queryPhoto(id int) bool {

	return true
}
