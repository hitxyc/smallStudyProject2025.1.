package mapper

import (
	"your_project/entity"

	"gorm.io/gorm"
)

// UserMapper 封装用户相关数据库操作
type UserMapper struct{}

// InsertUser 插入用户信息
func (um *UserMapper) InsertUser(db *gorm.DB, userName, password string) (int, error) {
	user := entity.UserEntity{
		UserName: userName,
		Password: password,
	}
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.UID, nil
}

// GetUid 根据用户名和密码查询 UID
func (um *UserMapper) GetUid(db *gorm.DB, password, userName string) (int, error) {
	var user entity.UserEntity
	if err := db.Where("password = ? AND user_name = ?", password, userName).First(&user).Error; err != nil {
		return 0, err
	}
	return user.UID, nil
}

// LoginByUid 根据 UID 和密码进行登录验证
func (um *UserMapper) LoginByUid(db *gorm.DB, uid int, password string) (bool, error) {
	var count int64
	if err := db.Model(&entity.UserEntity{}).
		Where("uid = ? AND password = ?", uid, password).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// LoginByUserName 根据用户名和密码进行登录验证
func (um *UserMapper) LoginByUserName(db *gorm.DB, userName, password string) (bool, error) {
	var count int64
	if err := db.Model(&entity.UserEntity{}).
		Where("user_name = ? AND password = ?", userName, password).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetUsernameByUid 根据 UID 查询用户名
func (um *UserMapper) GetUsernameByUid(db *gorm.DB, uid int) (string, error) {
	var user entity.UserEntity
	if err := db.Select("user_name").Where("uid = ?", uid).First(&user).Error; err != nil {
		return "", err
	}
	return user.UserName, nil
}

// GetUidRandom 随机返回一个 UID（排除指定的 UID）
func (um *UserMapper) GetUidRandom(db *gorm.DB, senderUid int) (int, error) {
	var user entity.UserEntity
	if err := db.Where("uid != ?", senderUid).Order("RAND()").Limit(1).First(&user).Error; err != nil {
		return 0, err
	}
	return user.UID, nil
}
