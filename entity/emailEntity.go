package entity

import "time"

// EmailEntity 定义邮件实体
type EmailEntity struct {
	ID          int       `json:"id" gorm:"column:id;primaryKey"`          // 邮件 ID
	SenderUid   int       `json:"sender_uid" gorm:"column:sender_uid"`     // 发件人 UID
	ReceiverUid int       `json:"receiver_uid" gorm:"column:receiver_uid"` // 收件人 UID
	Title       string    `json:"title" gorm:"column:title"`               // 邮件标题
	Content     string    `json:"content" gorm:"column:content"`           // 邮件内容
	SendTime    time.Time `json:"send_time" gorm:"column:send_time"`       // 发送时间
	IsRead      int       `json:"is_read" gorm:"column:is_read"`           // 是否已读（0: 未读, 1: 已读）
	GroupID     int       `json:"group_id" gorm:"column:group_id"`         // 群组 ID
	CanSee      int       `json:"can_see" gorm:"column:can_see"`           // 是否可见（0: 可见, 1: 不可见）
}

// TableName 指定表名
func (EmailEntity) TableName() string {
	return "email_entity"
}
