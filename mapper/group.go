package mapper

import (
	"your_project/entity"

	"gorm.io/gorm"
)

// GroupMapper 提供群组相关的数据库操作
type GroupMapper struct{}

// SendEmailToAll 根据用户 UID 和群组名称获取群组中的好友 UID
func (gm *GroupMapper) SendEmailToAll(db *gorm.DB, userUid int, groupName string) ([]int, error) {
	var friendUids []int
	err := db.Model(&entity.GroupEntity{}).
		Select("friend_uid").
		Where("user_uid = ? AND group_name = ?", userUid, groupName).
		Find(&friendUids).Error
	return friendUids, err
}
