package service

import (
	"time"
	"your_project/entity"
	"your_project/mapper"
)

type EmailService struct {
	EmailMapper *mapper.EmailMapper
}

func (s *EmailService) SendEmail(senderUid, receiverUid int, title, content string) (string, bool) {
	email := entity.EmailEntity{
		SenderUid:   senderUid,
		ReceiverUid: receiverUid,
		Title:       title,
		Content:     content,
		SendTime:    time.Now(),
	}
	err := s.EmailMapper.Insert(email)
	if err != nil {
		return "发送信件错误", false
	}
	return "发送成功", true
}

func (s *EmailService) ReceiveEmail(receiverUid, pageNum int) ([]entity.EmailEntity, string, bool) {
	pageSize := 5
	offset := (pageNum - 1) * pageSize
	emails, err := s.EmailMapper.ShowAllEmailEntity(receiverUid, offset, pageSize)
	if err != nil {
		return nil, "接收信件错误", false
	}
	return emails, "收件箱：", true
}

func (s *EmailService) SelectEmail(receiverUid int, keyword string) ([]entity.EmailEntity, string, bool) {
	emails, err := s.EmailMapper.SelectEmail(receiverUid, keyword)
	if err != nil {
		return nil, "模糊查询错误", false
	}
	return emails, "查询结果", true
}

func (s *EmailService) DriftEmail(senderUid int, title, content string) (string, bool) {
	email := entity.EmailEntity{
		SenderUid:   senderUid,
		ReceiverUid: 0, // Simulate random user for drift bottle
		Title:       title,
		Content:     content,
		SendTime:    time.Now(),
	}
	err := s.EmailMapper.Insert(email)
	if err != nil {
		return "漂流瓶错误", false
	}
	return "发送成功", true
}

func (s *EmailService) TimeEmail(senderUid, receiverUid int, title, content string, timeDelay int) (string, bool) {
	email := entity.EmailEntity{
		SenderUid:   senderUid,
		ReceiverUid: receiverUid,
		Title:       title,
		Content:     content,
		SendTime:    time.Now(),
		CanSee:      1, // Set as not visible initially
	}
	err := s.EmailMapper.Insert(email)
	if err != nil {
		return "定时发送错误", false
	}

	go func() {
		time.Sleep(time.Duration(timeDelay) * time.Minute)
		s.EmailMapper.UpdateCanSeeAndSendTime(time.Now(), senderUid, email.SendTime)
	}()

	return "发送成功", true
}

func (s *EmailService) DeleteEmail(senderUid, receiverUid int, title string) (string, bool) {
	// Implement delete functionality with appropriate mapping
	return "删除成功", true
}
