package entity

// ListEntity 定义好友列表实体
type ListEntity struct {
	FriendUid    int    `json:"friend_uid" gorm:"column:friend_uid"`     // 好友 UID
	UserName     string `json:"username" gorm:"column:username"`         // 用户名
	Remark       string `json:"remark" gorm:"column:remark"`             // 好友备注
	Relationship string `json:"relationship" gorm:"column:relationship"` // 亲密关系
	Status       int    `json:"status" gorm:"column:status"`             // 状态（0: 正常, 1: 假删除）
	UserUid      int    `json:"user_uid" gorm:"column:user_uid"`         // 用户 UID
	IsFriend     int    `json:"is_friend" gorm:"column:is_friend"`       // 是否是好友（0: 不是, 1: 是）
}

// TableName 指定表名
func (ListEntity) TableName() string {
	return "list_entity"
}
