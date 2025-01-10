package service

import (
	"errors"
	"your_project/entity"
	"your_project/mapper"
	"your_project/utils"

	"gorm.io/gorm"
)

type ListService struct {
	ListMapper *mapper.ListMapper
	UserMapper *mapper.UserMapper
}

// InsertFriend 添加好友
func (ls *ListService) InsertFriend(db *gorm.DB, friendUid int, userUid int) (*entity.ResultEntity, error) {
	userName, err := ls.UserMapper.GetUsernameByUid(db, friendUid)
	if err != nil {
		return nil, errors.New("朋友 UID 错误")
	}

	listEntity := entity.ListEntity{
		FriendUid: friendUid,
		UserName:  userName,
		UserUid:   userUid,
	}
	if userName == "" {
		return utils.NewResultEntity("UID 错误", false, nil), nil
	}

	if err := db.Create(&listEntity).Error; err != nil {
		return nil, err
	}
	return utils.NewResultEntity("添加成功", true, nil), nil
}

// UpdateRemark 更新好友备注
func (ls *ListService) UpdateRemark(db *gorm.DB, userUid int, friendUid int, remark string) (*entity.ResultEntity, error) {
	if err := ls.ListMapper.UpdateRemark(db, userUid, friendUid, remark); err != nil {
		return nil, errors.New("添加备注错误")
	}
	return utils.NewResultEntity("添加备注成功", true, nil), nil
}

// SetRelationship 设置亲密关系
func (ls *ListService) SetRelationship(db *gorm.DB, userUid int, friendUid int, relationship string) (*entity.ResultEntity, error) {
	if err := ls.ListMapper.SetRelationship(db, userUid, friendUid, relationship); err != nil {
		return nil, errors.New("设置亲密关系错误")
	}
	return utils.NewResultEntity("设置亲密关系成功", true, nil), nil
}

// DeleteFriend 假删除好友
func (ls *ListService) DeleteFriend(db *gorm.DB, userUid int, friendUid int) (*entity.ResultEntity, error) {
	if err := ls.ListMapper.DeleteFriend(db, userUid, friendUid); err != nil {
		return nil, errors.New("假删除错误")
	}
	return utils.NewResultEntity("成功假删除好友", true, nil), nil
}

// CheckWhoAdd 查看好友申请
func (ls *ListService) CheckWhoAdd(db *gorm.DB, friendUid int) (*entity.ResultEntity, error) {
	userUids, err := ls.ListMapper.CheckWhoAdd(db, friendUid)
	if err != nil {
		return nil, errors.New("好友 UID 错误")
	}
	return utils.NewResultEntity("好友申请", true, userUids), nil
}

// AgreeFriendAdd 同意好友申请
func (ls *ListService) AgreeFriendAdd(db *gorm.DB, userUid int, friendUid int) (*entity.ResultEntity, error) {
	if err := ls.ListMapper.AgreeFriendAddUser(db, userUid, friendUid); err != nil {
		return nil, errors.New("同意好友申请错误")
	}

	userName, err := ls.UserMapper.GetUsernameByUid(db, friendUid)
	if err != nil {
		return nil, errors.New("获取好友用户名错误")
	}

	listEntity := entity.ListEntity{
		FriendUid: friendUid,
		UserName:  userName,
		UserUid:   userUid,
		IsFriend:  1,
	}
	if err := db.Create(&listEntity).Error; err != nil {
		return nil, err
	}
	return utils.NewResultEntity("同意好友申请成功", true, nil), nil
}

// ShowAllFriend 查看好友列表
func (ls *ListService) ShowAllFriend(db *gorm.DB, userUid int, pageNum int) (*entity.ResultEntity, error) {
	pageSize := 10
	offset := (pageNum - 1) * pageSize
	friends, err := ls.ListMapper.ShowAllFriend(db, userUid, pageSize, offset)
	if err != nil {
		return nil, errors.New("好友列表展示错误")
	}
	return utils.NewResultEntity("好友列表", true, friends), nil
}
