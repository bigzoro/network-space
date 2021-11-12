package dao

import (
	"personSpace/model"
)

type DocumentDao struct {
}

// 添加文件
func (dd *DocumentDao) addDocument(Document model.Document) bool {

	return true
}

// 删除文件
func (dd *DocumentDao) delteteDocument(id int) bool {

	return true
}

// 修改文件
func (dd *DocumentDao) updateDocument(id int, Document model.Document) bool {

	return true
}

// 查询文件
func (dd *DocumentDao) queryDocument(id int) bool {

	return true
}
