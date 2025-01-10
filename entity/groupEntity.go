package entity

// GroupEntity 定义群组实体
type GroupEntity struct {
	UserUid   int    `json:"user_uid" gorm:"column:user_uid"`     // 用户 UID
	FriendUid int    `json:"friend_uid" gorm:"column:friend_uid"` // 好友 UID
	GroupName string `json:"group_name" gorm:"column:group_name"` // 群组名称
	Status    int    `json:"status" gorm:"column:status"`         // 群组状态（0: 正常, 1: 删除）
}

// TableName 指定表名
func (GroupEntity) TableName() string {
	return "group_entity"
}
