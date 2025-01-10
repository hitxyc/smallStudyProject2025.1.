package service

import (
	"github.com/dgrijalva/jwt-go"

	"gorm.io/gorm"
	"your_project/entity"
	"your_project/mapper"
	"your_project/utils"
)

type LoginService struct {
	UserMapper *mapper.UserMapper
}

// SaveUser 注册用户
func (ls *LoginService) SaveUser(db *gorm.DB, userName string, password string) (*entity.ResultEntity, error) {
	resultEntity := &entity.ResultEntity{}

	// 密码长度校验
	if len(password) < 6 {
		return resultEntity.Construct("密码长度不得小于6位", false, nil), nil
	}

	// 验证用户名和密码是否重复
	md5Password := utils.MD5(password)
	exists, err := ls.UserMapper.LoginByUserName(db, userName, md5Password)
	if err != nil {
		return nil, err
	}
	if exists {
		return resultEntity.Construct("用户名密码重复", false, nil), nil
	}

	// 插入用户数据
	uid, err := ls.UserMapper.InsertUser(db, userName, md5Password)
	if err != nil {
		return nil, err
	}

	return resultEntity.Construct("注册成功，UID:", true, uid), nil
}

// Login 用户登录
func (ls *LoginService) Login(db *gorm.DB, userName string, password string, uid *int) (*entity.ResultEntity, error) {
	resultEntity := &entity.ResultEntity{}
	md5Password := utils.MD5(password)

	// 用户名密码登录
	exists, err := ls.UserMapper.LoginByUserName(db, userName, md5Password)
	if err != nil {
		return nil, err
	}
	if exists {
		claims := jwt.MapClaims{
			"name": userName,
		}
		token, err := utils.GenerateJwt(claims)
		if err != nil {
			return nil, err
		}
		userUid, err := ls.UserMapper.GetUid(db, md5Password, userName)
		if err != nil {
			return nil, err
		}
		return resultEntity.Construct("登录成功", true, map[string]interface{}{
			"uid": userUid,
			"jwt": token,
		}), nil
	}

	// UID 密码登录
	if uid != nil {
		exists, err = ls.UserMapper.LoginByUid(db, *uid, md5Password)
		if err != nil {
			return nil, err
		}
		if exists {
			claims := jwt.MapClaims{
				"uid": *uid,
			}
			token, err := utils.GenerateJwt(claims)
			if err != nil {
				return nil, err
			}
			return resultEntity.Construct("登录成功", true, map[string]interface{}{
				"uid": *uid,
				"jwt": token,
			}), nil
		}
	}

	return resultEntity.Construct("账号或密码错误", false, nil), nil
}
