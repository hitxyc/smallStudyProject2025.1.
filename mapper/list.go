package mapper

import (
	"your_project/entity"

	"gorm.io/gorm"
)

// ListMapper 提供好友列表相关的数据库操作
type ListMapper struct{}

// UpdateRemark 更新好友备注
func (lm *ListMapper) UpdateRemark(db *gorm.DB, userUid int, friendUid int, remark string) error {
	return db.Model(&entity.ListEntity{}).
		Where("user_uid = ? AND friend_uid = ?", userUid, friendUid).
		Update("remark", remark).Error
}

// SetRelationship 设置亲密关系
func (lm *ListMapper) SetRelationship(db *gorm.DB, userUid int, friendUid int, relationship string) error {
	return db.Model(&entity.ListEntity{}).
		Where("user_uid = ? AND friend_uid = ?", userUid, friendUid).
		Update("relationship", relationship).Error
}

// DeleteFriend 假删除好友
func (lm *ListMapper) DeleteFriend(db *gorm.DB, userUid int, friendUid int) error {
	return db.Model(&entity.ListEntity{}).
		Where("user_uid = ? AND friend_uid = ?", userUid, friendUid).
		Update("status", 1).Error
}

// CheckWhoAdd 查询好友申请
func (lm *ListMapper) CheckWhoAdd(db *gorm.DB, friendUid int) ([]int, error) {
	var userUids []int
	err := db.Model(&entity.ListEntity{}).
		Select("user_uid").
		Where("friend_uid = ? AND is_friend = 0", friendUid).
		Find(&userUids).Error
	return userUids, err
}

// AgreeFriendAddUser 同意好友申请
func (lm *ListMapper) AgreeFriendAddUser(db *gorm.DB, userUid int, friendUid int) error {
	return db.Model(&entity.ListEntity{}).
		Where("user_uid = ? AND friend_uid = ? AND is_friend = 0", userUid, friendUid).
		Update("is_friend", 1).Error
}

// ShowAllFriend 分页查看好友列表
func (lm *ListMapper) ShowAllFriend(db *gorm.DB, userUid int, limit int, offset int) ([]entity.ListEntity, error) {
	var friends []entity.ListEntity
	err := db.Where("user_uid = ? AND is_friend = 1 AND status = 0", userUid).
		Limit(limit).
		Offset(offset).
		Find(&friends).Error
	return friends, err
}
