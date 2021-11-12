package dao

import (
	"personSpace/model"
)

type FriendDao struct {
}

// 添加朋友
func (dd *FriendDao) addFriend(Friend model.Friend) bool {

	return true
}

// 删除朋友
func (dd *FriendDao) delteteFriend(id int) bool {

	return true
}

// 修改朋友
func (dd *FriendDao) updateFriend(id int, Friend model.Friend) bool {

	return true
}

// 查询朋友
func (dd *FriendDao) queryFriend(id int) bool {

	return true
}
