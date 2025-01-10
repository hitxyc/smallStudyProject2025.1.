package service

import (
	"errors"
	"time"
	"your_project/entity"
	"your_project/mapper"
	"your_project/utils"

	"gorm.io/gorm"
)

type GroupService struct {
	GroupMapper *mapper.GroupMapper
	EmailMapper *mapper.EmailMapper
}

// InsertGroup 创建群组
func (gs *GroupService) InsertGroup(db *gorm.DB, userUid int, friendUid int, groupName string) (*entity.ResultEntity, error) {
	groupEntity := entity.GroupEntity{
		UserUid:   userUid,
		FriendUid: friendUid,
		GroupName: groupName,
	}
	if err := db.Create(&groupEntity).Error; err != nil {
		return nil, errors.New("添加群组错误")
	}
	return utils.NewResultEntity("添加小组成功", true, nil), nil
}

// SendEmailToAll 群发消息
func (gs *GroupService) SendEmailToAll(db *gorm.DB, userUid int, groupName string, title string, content string) (*entity.ResultEntity, error) {
	// 获取当前时间
	sendTime := time.Now()

	// 获取群组中的好友 UID
	friendUids, err := gs.GroupMapper.SendEmailToAll(db, userUid, groupName)
	if err != nil {
		return nil, errors.New("获取群组成员失败")
	}

	// 为群组中的每个好友发送邮件
	for _, friendUid := range friendUids {
		emailEntity := entity.EmailEntity{
			SenderUid:   userUid,
			ReceiverUid: friendUid,
			Title:       title,
			Content:     content,
			SendTime:    sendTime,
		}
		if err := db.Create(&emailEntity).Error; err != nil {
			return nil, errors.New("发送邮件失败")
		}
	}

	return utils.NewResultEntity("发送成功", true, nil), nil
}
