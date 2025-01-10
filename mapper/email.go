package mapper

import (
	"gorm.io/gorm"
	"time"
	"your_project/entity"
)

type EmailMapper struct {
	DB *gorm.DB
}

func (e *EmailMapper) UpdateIsReadByReceiverUid(receiverUid int) error {
	return e.DB.Model(&entity.EmailEntity{}).Where("receiver_uid = ?", receiverUid).Update("is_read", 1).Error
}

func (e *EmailMapper) ShowAllEmailEntity(receiverUid, offset, limit int) ([]entity.EmailEntity, error) {
	var emails []entity.EmailEntity
	err := e.DB.Where("receiver_uid = ? OR sender_uid = ? AND can_see = 0", receiverUid, receiverUid).
		Order("send_time DESC").Offset(offset).Limit(limit).Find(&emails).Error
	return emails, err
}

func (e *EmailMapper) SelectEmail(receiverUid int, keyword string) ([]entity.EmailEntity, error) {
	var emails []entity.EmailEntity
	err := e.DB.Where("receiver_uid = ? OR sender_uid = ? AND can_see = 0 AND (title LIKE ? OR content LIKE ?)",
		receiverUid, receiverUid, "%"+keyword+"%", "%"+keyword+"%").
		Order("send_time DESC").Find(&emails).Error
	return emails, err
}

func (e *EmailMapper) Insert(email entity.EmailEntity) error {
	return e.DB.Create(&email).Error
}

func (e *EmailMapper) UpdateCanSeeAndSendTime(newSendTime time.Time, senderUid int, oldSendTime time.Time) error {
	return e.DB.Model(&entity.EmailEntity{}).Where("sender_uid = ? AND send_time = ?", senderUid, oldSendTime).
		Updates(map[string]interface{}{"send_time": newSendTime, "can_see": 0}).Error
}
