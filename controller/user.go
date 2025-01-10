package controller

import (
	"gorm.io/gorm"
	"net/http"
	"your_project/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService *service.LoginService
}

// SaveUser 注册用户接口
// @Summary 注册用户
// @Description 注册新用户，返回用户的 UID
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param userName body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} entity.ResultEntity
// @Router /login/saveUser [post]
func (lc *LoginController) SaveUser(c *gin.Context) {
	var request struct {
		UserName string `json:"userName" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	result, err := lc.LoginService.SaveUser(c.MustGet("db").(*gorm.DB), request.UserName, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Login 用户登录接口
// @Summary 用户登录
// @Description 用户通过用户名/UID 和密码登录，返回 JWT Token
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param userName body string false "用户名"
// @Param password body string true "密码"
// @Param uid body int false "用户 ID"
// @Success 200 {object} entity.ResultEntity
// @Router /login/log [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request struct {
		UserName string `json:"userName"`
		Password string `json:"password" binding:"required"`
		UID      *int   `json:"uid"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	result, err := lc.LoginService.Login(c.MustGet("db").(*gorm.DB), request.UserName, request.Password, request.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
